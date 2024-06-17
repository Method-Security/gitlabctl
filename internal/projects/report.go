package projects

import (
	"github.com/xanzy/go-gitlab"
)

// GitlabResources represents a collection of Gitlab projects.
type GitlabResources struct {
	Projects []*gitlab.Project `json:"projects" yaml:"projects"`
}

// GitlabResourceReport represents a report of Gitlab resources and non-fatal errors encountered during enumeration.
type GitlabResourceReport struct {
	BaseURL   string          `json:"base_url" yaml:"base_url"`
	Resources GitlabResources `json:"resources" yaml:"resources"`
	Errors    []string        `json:"errors" yaml:"errors"`
}
