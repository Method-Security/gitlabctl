// Package projects holds the data structures and logic necessary to interact with the Gitlab API and enumerate projects.
package projects

import (
	"context"
	"fmt"

	"github.com/xanzy/go-gitlab"
)

// EnumerateProjectsOptions holds the options for enumerating projects.
// The Mine field is used to filter projects that are owned by the authenticated user, only returning projects that are owned
// by the authenticated user when set to true.
// The Archived field is used to filter for archived projects, including archived when set to true.
// The GroupID field is used to filter projects by group ID, only returning projects that are part of the specified group.
type EnumerateProjectsOptions struct {
	Mine     bool   `json:"mine"`
	Archived bool   `json:"archived"`
	GroupID  string `json:"group_id"`
}

// FindGroupByName searches for a group by name using the provided Gitlab client. If the group is found, it is returned.
// If the group is not found, an error is returned.
func FindGroupByName(ctx context.Context, client *gitlab.Client, groupName string) (*gitlab.Group, error) {
	options := &gitlab.ListGroupsOptions{
		Search: gitlab.Ptr(groupName),
	}

	groups, _, err := client.Groups.ListGroups(options)
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		if group.Name == groupName {
			return group, nil
		}
	}

	return nil, fmt.Errorf("group %s not found", groupName)
}

// EnumerateProjects enumerates projects using the provided Gitlab client and options. The function returns a GitlabResourceReport
// containing the resources and non-fatal errors encountered during the enumeration process.
func EnumerateProjects(ctx context.Context, baseURL string, options *EnumerateProjectsOptions, client *gitlab.Client) (*GitlabResourceReport, error) {
	report := &GitlabResourceReport{
		Resources: GitlabResources{},
		Errors:    []string{},
		BaseURL:   baseURL,
	}
	filterOptions := gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	if options.Archived {
		filterOptions.Archived = gitlab.Ptr(options.Archived)
	}
	if options.Mine {
		filterOptions.Owned = gitlab.Ptr(options.Mine)
	}

	for {
		projects, resp, err := client.Projects.ListProjects(&filterOptions)
		if err != nil {
			report.Errors = append(report.Errors, err.Error())
			break
		}

		report.Resources.Projects = append(report.Resources.Projects, projects...)
		if resp.NextPage == 0 {
			break
		}
		filterOptions.ListOptions.Page = resp.NextPage
	}

	return report, nil
}

// EnumerateProjectsForGroup enumerates projects for a specific group using the provided Gitlab client and options. The function
// returns a GitlabResourceReport containing the resources and non-fatal errors encountered during the enumeration process.
func EnumerateProjectsForGroup(ctx context.Context, baseURL string, client *gitlab.Client, options *EnumerateProjectsOptions) (*GitlabResourceReport, error) {
	report := &GitlabResourceReport{
		Resources: GitlabResources{},
		Errors:    []string{},
		BaseURL:   baseURL,
	}

	err := fetchGroupAndSubgroupProjects(ctx, client, options.GroupID, options, report)
	if err != nil {
		report.Errors = append(report.Errors, err.Error())
	}

	return report, nil
}

func fetchGroupAndSubgroupProjects(ctx context.Context, client *gitlab.Client, groupID string, options *EnumerateProjectsOptions, report *GitlabResourceReport) error {
	filterOptions := gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	if options.Archived {
		filterOptions.Archived = gitlab.Ptr(options.Archived)
	}
	if options.Mine {
		filterOptions.Owned = gitlab.Ptr(options.Mine)
	}

	for {
		projects, resp, err := client.Groups.ListGroupProjects(groupID, &filterOptions)
		if err != nil {
			return err
		}

		report.Resources.Projects = append(report.Resources.Projects, projects...)
		if resp.NextPage == 0 {
			break
		}
		filterOptions.ListOptions.Page = resp.NextPage
	}

	// Fetch subgroups
	subGroupOptions := gitlab.ListSubGroupsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}

	for {
		subgroups, resp, err := client.Groups.ListSubGroups(groupID, &subGroupOptions)
		if err != nil {
			return err
		}

		for _, subgroup := range subgroups {
			err := fetchGroupAndSubgroupProjects(ctx, client, fmt.Sprintf("%d", subgroup.ID), options, report)
			if err != nil {
				report.Errors = append(report.Errors, err.Error())
			}
		}

		if resp.NextPage == 0 {
			break
		}
		subGroupOptions.Page = resp.NextPage
	}

	return nil
}
