package cmd

import (
	"github.com/Method-Security/gitlabctl/internal/projects"
	"github.com/spf13/cobra"
)

// InitProjectsCmd initializes the projects command for the gitlabctl CLI. This command sets up the flags for the command,
// parsing the provided group ID, archived, and mine flags before passing them to the projects package for enumeration.
func (a *Gitlabctl) InitProjectsCmd() {
	options := projects.EnumerateProjectsOptions{
		Mine:     true,
		Archived: false,
		GroupID:  "",
	}

	a.ProjectsCmd = &cobra.Command{
		Use:   "projects",
		Short: "Enumerate Gitlab projects",
		Long:  `Enumerate Gitlab projects`,
		Run: func(cmd *cobra.Command, args []string) {
			var report *projects.GitlabResourceReport
			var err error
			if options.GroupID == "" {
				report, err = projects.EnumerateProjects(cmd.Context(), a.RootFlags.BaseURL, &options, a.GitlabClient)
			} else {
				report, err = projects.EnumerateProjectsForGroup(cmd.Context(), a.RootFlags.BaseURL, a.GitlabClient, &options)
			}
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report
		},
	}
	a.ProjectsCmd.Flags().BoolVar(&options.Archived, "archived", false, "Include archived projects")
	a.ProjectsCmd.Flags().BoolVar(&options.Mine, "mine", true, "Include only projects owned by the authenticated user.")
	a.ProjectsCmd.Flags().StringVar(&options.GroupID, "group-id", "", "Group ID")

	a.RootCmd.AddCommand(a.ProjectsCmd)
}
