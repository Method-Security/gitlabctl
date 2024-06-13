package main

import (
	"flag"
	"os"

	"github.com/Method-Security/gitlabctl/cmd"
)

var version = "none"

func main() {
	flag.Parse()

	gitlabctl := cmd.NewGitlabctl(version)
	gitlabctl.InitRootCommand()
	gitlabctl.InitProjectsCmd()
	gitlabctl.InitVulnerabilityCmd()

	if err := gitlabctl.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
