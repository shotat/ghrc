package spec

import (
	"github.com/shotat/ghrc/change"
	"github.com/shotat/ghrc/state"
)

type Repo struct {
	Description      *string `yaml:"description,omitempty"`
	Homepage         *string `yaml:"homepage,omitempty"`
	Private          *bool   `yaml:"private"`
	AllowSquashMerge *bool   `yaml:"allowSquashMerge"`
	AllowMergeCommit *bool   `yaml:"allowMergeCommit"`
	AllowRebaseMerge *bool   `yaml:"allowRebaseMerge"`

	Topics []string `yaml:"topics,omitempty"`
}

func LoadRepoSpecFromState(st *state.Repo) *Repo {
	return &Repo{
		Homepage:         st.Homepage,
		Description:      st.Description,
		Private:          st.Private,
		Topics:           st.Topics,
		AllowSquashMerge: st.AllowSquashMerge,
		AllowMergeCommit: st.AllowMergeCommit,
		AllowRebaseMerge: st.AllowRebaseMerge,
	}
}

func (sp *Repo) GetRepoChange(st *state.Repo) *change.ReposChange {
	after := &state.Repo{
		ID:               st.ID,
		Name:             st.Name,
		Owner:            st.Owner,
		Description:      st.Description,
		Homepage:         st.Homepage,
		Private:          st.Private,
		AllowSquashMerge: st.AllowSquashMerge,
		AllowMergeCommit: st.AllowMergeCommit,
		AllowRebaseMerge: st.AllowRebaseMerge,
		Topics:           st.Topics,
	}

	if sp.Description != nil {
		after.Description = sp.Description
	}
	if sp.Private != nil {
		after.Private = sp.Private
	}
	if sp.Homepage != nil {
		after.Homepage = sp.Homepage
	}
	if sp.AllowSquashMerge != nil {
		after.AllowSquashMerge = sp.AllowSquashMerge
	}
	if sp.AllowMergeCommit != nil {
		after.AllowMergeCommit = sp.AllowMergeCommit
	}
	if sp.AllowRebaseMerge != nil {
		after.AllowRebaseMerge = sp.AllowRebaseMerge
	}
	if sp.Topics != nil {
		after.Topics = sp.Topics
	}
	return &change.ReposChange{
		Action: change.Update,
		Before: st,
		After:  after,
	}
}
