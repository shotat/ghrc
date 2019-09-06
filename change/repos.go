package change

import (
	"github.com/shotat/ghrc/spec"
	"github.com/shotat/ghrc/state"
)

func GetRepoChange(st *state.Repo, sp *spec.Repo) *ReposChange {
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
	return &ReposChange{
		Action: Update,
		Before: st,
		After:  after,
	}
}
