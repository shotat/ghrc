package change

import (
	"bytes"
	"context"
	"fmt"
	"sort"

	"github.com/google/go-cmp/cmp"
	"github.com/shotat/ghrc/spec"
	"github.com/shotat/ghrc/state"
)

type RepoChange struct {
	Action Action
	Before *state.Repo
	After  *state.Repo
}

func (c *RepoChange) Apply(ctx context.Context, repoOwner string, repoName string) error {
	return c.After.Update(ctx, repoOwner, repoName)
}

func (c *RepoChange) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(c.subject())

	// This Transformer sorts a []int.
	trans := cmp.Transformer("Sort", func(in []string) []string {
		out := append([]string(nil), in...) // Copy input to avoid mutating it
		sort.Strings(out)
		return out
	})
	diff := cmp.Diff(c.Before, c.After, trans)
	buf.WriteString(diff)
	return buf.String()
}

func (c *RepoChange) subject() string {
	var name string
	switch c.Action {
	case Create, Update:
		name = fmt.Sprintf("%s/%s", c.After.Owner, c.After.Name)
	case Delete:
		name = fmt.Sprintf("%s/%s", c.Before.Owner, c.Before.Name)
	}
	return fmt.Sprintf("%s Repo: %s\n", string(c.Action), name)
}

func GetRepoChange(st *state.Repo, sp *spec.Repo) *RepoChange {
	if sp == nil {
		return nil
	}

	after := &state.Repo{
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
		after.Description = *sp.Description
	}
	if sp.Homepage != nil {
		after.Homepage = *sp.Homepage
	}
	if sp.Private != nil {
		after.Private = *sp.Private
	}
	if sp.AllowSquashMerge != nil {
		after.AllowSquashMerge = *sp.AllowSquashMerge
	}
	if sp.AllowMergeCommit != nil {
		after.AllowMergeCommit = *sp.AllowMergeCommit
	}
	if sp.AllowRebaseMerge != nil {
		after.AllowRebaseMerge = *sp.AllowRebaseMerge
	}
	if sp.Topics != nil {
		after.Topics = sp.Topics
	}
	return &RepoChange{
		Action: Update,
		Before: st,
		After:  after,
	}
}
