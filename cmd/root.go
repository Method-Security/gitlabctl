// Package cmd implements the CobraCLI commands for the gitlabctl CLI. Subcommands for the CLI should all live within
// this package. Logic should be delegated to internal packages and functions to keep the CLI commands clean and
// focused on CLI I/O.
package cmd

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/Method-Security/gitlabctl/internal/config"
	"github.com/Method-Security/pkg/signal"
	"github.com/Method-Security/pkg/writer"
	"github.com/palantir/pkg/datetime"
	"github.com/palantir/witchcraft-go-logging/wlog/svclog/svc1log"
	"github.com/spf13/cobra"
	"github.com/xanzy/go-gitlab"
)

// Gitlabctl is the main struct for the gitlabctl CLI. It contains the version, root flags, output config, output signal,
// information, providing a context for subcommands to leverage during execution. The output signal is used to write the
// output of the command to the desired output format and location.
type Gitlabctl struct {
	Version          string
	RootFlags        config.RootFlags
	OutputConfig     writer.OutputConfig
	OutputSignal     signal.Signal
	RootCmd          *cobra.Command
	VersionCmd       *cobra.Command
	ProjectsCmd      *cobra.Command
	VulnerabilityCmd *cobra.Command
	GitlabClient     *gitlab.Client
}

// NewGitlabctl creates a new Gitlabctl struct with the provided version. The root flags, output config, and output format.
// We pass the version in here from our main.go file, where we set the version string during the build process.
func NewGitlabctl(version string) *Gitlabctl {
	gitlabctl := Gitlabctl{
		Version: version,
		RootFlags: config.RootFlags{
			Quiet:   false,
			Verbose: false,
			BaseURL: "",
			Token:   "",
		},
		OutputConfig: writer.NewOutputConfig(nil, writer.NewFormat(writer.SIGNAL)),
		OutputSignal: signal.NewSignal(nil, datetime.DateTime(time.Now()), nil, 0, nil),
	}
	return &gitlabctl
}

// InitRootCommand initializes the root command for the gitlabctl CLI. This command sets up the persistent flags for the
// CLI, including the quiet, verbose, base-url, token, output-file, and output flags. The root command also sets up the
// version command, which prints the version of the gitlabctl CLI.
// The root command sets the PersistentPreRunE, which is responsible for initializing the output signal, as well as creating
// the Gitlab client that will be used in all commands. The PersistentPostRunE is responsible for writing the output of the
// command to the desired output format and location.
func (a *Gitlabctl) InitRootCommand() {
	var outputFormat string
	var outputFile string
	a.RootCmd = &cobra.Command{
		Use:   "gitlabctl",
		Short: "Gitlabctl CLI",
		Long:  `Gitlabctl CLI`,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			cmd.SetContext(svc1log.WithLogger(cmd.Context(), config.InitializeLogging(cmd, &a.RootFlags)))
			var token string
			if os.Getenv("GITLAB_TOKEN") != "" {
				token = os.Getenv("GITLAB_TOKEN")
			} else if a.RootFlags.Token != "" {
				token = a.RootFlags.Token
			} else {
				return errors.New("either GITLAB_TOKEN environment variable or --token must be set")
			}
			a.GitlabClient, _ = gitlab.NewClient(token, gitlab.WithBaseURL(a.RootFlags.BaseURL))

			if a.RootFlags.BaseURL == "" {
				return errors.New("base-url flag not set")
			}
			a.RootFlags.BaseURL = config.NormalizeGitlabURL(a.RootFlags.BaseURL)

			format, err := validateOutputFormat(outputFormat)
			if err != nil {
				return err
			}
			var outputFilePointer *string
			if outputFile != "" {
				outputFilePointer = &outputFile
			} else {
				outputFilePointer = nil
			}
			a.OutputConfig = writer.NewOutputConfig(outputFilePointer, format)
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			completedAt := datetime.DateTime(time.Now())
			a.OutputSignal.CompletedAt = &completedAt
			return writer.Write(
				a.OutputSignal.Content,
				a.OutputConfig,
				a.OutputSignal.StartedAt,
				a.OutputSignal.CompletedAt,
				a.OutputSignal.Status,
				a.OutputSignal.ErrorMessage,
			)
		},
	}

	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.Quiet, "quiet", "q", false, "Suppress output")
	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.Verbose, "verbose", "v", false, "Verbose output")
	a.RootCmd.PersistentFlags().StringVar(&a.RootFlags.BaseURL, "base-url", "", "Base URL for Gitlab API. (e.g. https://gitlab.com/api/v4)")
	a.RootCmd.PersistentFlags().StringVar(&a.RootFlags.Token, "token", "", "Gitlab Access Token. Can also be set via GITLAB_TOKEN environment variable")
	a.RootCmd.PersistentFlags().StringVarP(&outputFile, "output-file", "f", "", "Path to output file. If blank, will output to STDOUT")
	a.RootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "signal", "Output format (signal, json, yaml). Default value is signal")

	a.VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version number of gitlabctl",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(a.Version)
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}
	a.RootCmd.AddCommand(a.VersionCmd)
}

func validateOutputFormat(output string) (writer.Format, error) {
	var format writer.FormatValue
	switch strings.ToLower(output) {
	case "json":
		format = writer.JSON
	case "yaml":
		format = writer.YAML
	case "signal":
		format = writer.SIGNAL
	default:
		return writer.Format{}, errors.New("invalid output format. Valid formats are: json, yaml, signal")
	}
	return writer.NewFormat(format), nil
}
