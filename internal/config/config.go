// Package config contains common configuration values that are used by the various commands and subcommands in the CLI.
package config

// The RootFlags struct contains the common flags that are used by the various commands and subcommands in the CLI.
type RootFlags struct {
	Quiet   bool
	Verbose bool
	BaseURL string
	Token   string
}
