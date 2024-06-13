package projects

import (
	"github.com/xanzy/go-gitlab"
)

type GitlabResources struct {
	Projects []*gitlab.Project `json:"projects" yaml:"projects"`
}

type GitlabResourceReport struct {
	BaseURL   string          `json:"base_url" yaml:"base_url"`
	Resources GitlabResources `json:"resources" yaml:"resources"`
	Errors    []string        `json:"errors" yaml:"errors"`
}
