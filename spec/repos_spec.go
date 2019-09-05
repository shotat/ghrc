package spec

import (
	"github.com/shotat/ghrc/status"
)

type RepositorySpec struct {
	Description      *string `yaml:"description,omitempty"`
	Homepage         *string `yaml:"homepage,omitempty"`
	Private          bool    `yaml:"private"`
	AllowSquashMerge bool    `yaml:"allowSquashMerge"`
	AllowMergeCommit bool    `yaml:"allowMergeCommit"`
	AllowRebaseMerge bool    `yaml:"allowRebaseMerge"`

	Topics      []string     `yaml:"topics,omitempty"`
	Labels      []Label      `yaml:"labels,omitempty"`
	Protections []Protection `yaml:"protections,omitempty"`
}

type Label struct {
	Name        string  `yaml:"name"`
	Description *string `yaml:"description,omitempty"`
	Color       string  `yaml:"color"`
}

func (rs *RepositorySpec) Patch(repo *status.RepositoryStatus) {
	repo.Description = rs.Description
	repo.Private = rs.Private
	repo.Homepage = rs.Homepage
	repo.AllowSquashMerge = rs.AllowSquashMerge
	repo.AllowMergeCommit = rs.AllowMergeCommit
	repo.AllowRebaseMerge = rs.AllowRebaseMerge
}
