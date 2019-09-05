package spec

import (
	"github.com/shotat/ghrc/status"
)

type RepositorySpec struct {
	Description      *string `yaml:"description,omitempty"`
	Homepage         *string `yaml:"homepage,omitempty"`
	Private          *bool   `yaml:"private"`
	AllowSquashMerge *bool   `yaml:"allowSquashMerge"`
	AllowMergeCommit *bool   `yaml:"allowMergeCommit"`
	AllowRebaseMerge *bool   `yaml:"allowRebaseMerge"`

	Topics      []string     `yaml:"topics,omitempty"`
	Labels      []Label      `yaml:"labels,omitempty"`
	Protections []Protection `yaml:"protections,omitempty"`
}

func (sp *RepositorySpec) Patch(st *status.RepositoryStatus) {
	if sp.Description != nil {
		st.Description = sp.Description
	}
	if sp.Private != nil {
		st.Private = sp.Private
	}
	if sp.Homepage != nil {
		st.Homepage = sp.Homepage
	}
	if sp.AllowSquashMerge != nil {
		st.AllowSquashMerge = sp.AllowSquashMerge
	}
	if sp.AllowMergeCommit != nil {
		st.AllowMergeCommit = sp.AllowMergeCommit
	}
	if sp.AllowRebaseMerge != nil {
		st.AllowRebaseMerge = sp.AllowRebaseMerge
	}
	if sp.Topics != nil {
		st.Topics = sp.Topics
	}
	/*
		if sp.Protections != nil {
			st.Protections = sp.Protections
		}
	*/
}

type Label struct {
	Name        string  `yaml:"name"`
	Description *string `yaml:"description,omitempty"`
	Color       string  `yaml:"color"`
}

type Protection struct {
	Branch                     *string                     `yaml:"branch"`
	RequiredStatusCheck        *RequiredStatusCheck        `yaml:"requiredStatusCheck"`
	EnforceAdmins              *bool                       `yaml:"enforceAdmins"`
	RequiredPullRequestReviews *RequiredPullRequestReviews `yaml:"requiredPullRequestReviews"`
	Restrictions               *Restrictions               `yaml:"restrictions"`
}

type RequiredPullRequestReviews struct {
	DismissalRestrictions        *Restrictions `yaml:"dismissalRestrictions"`
	DismissStaleReviews          bool          `yaml:"dismissStaleReviews"`
	RequireCodeOwnerReviews      bool          `yaml:"requireCodeOwnerReviews"`
	RequiredApprovingReviewCount int           `yaml:"requiredApprovingReviewCount"`
}

type Restrictions struct {
	Users []string `yaml:"users"`
	Teams []string `yaml:"teams"`
}

type RequiredStatusCheck struct {
	Strict   bool     `yaml:"strict"`
	Contexts []string `yaml:"contexts"`
}
