package projects

import (
	"context"
	"fmt"

	"github.com/xanzy/go-gitlab"
)

type EnumerateProjectsOptions struct {
	Mine     bool   `json:"mine"`
	Archived bool   `json:"archived"`
	GroupID  string `json:"group_id"`
}

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
