package change

import (
	"bytes"
	"context"
	"fmt"

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
	buf.WriteString(fmt.Sprintf("%s Repo\n", string(c.Action)))
	switch c.Action {
	case Update:
		buf.WriteString(fmt.Sprintf("\tdescription\t%v\n", *c.After.Description))
		buf.WriteString(fmt.Sprintf("\thomepage\t%v\n", *c.After.Homepage))
		buf.WriteString(fmt.Sprintf("\tprivate\t%v\n", *c.After.Private))
		buf.WriteString(fmt.Sprintf("\tallowSquashMerge\t%v\n", *c.After.AllowSquashMerge))
		buf.WriteString(fmt.Sprintf("\tallowMergeCommit\t%v\n", *c.After.AllowMergeCommit))
		buf.WriteString(fmt.Sprintf("\tallowRebaseMerge\t%v\n", *c.After.AllowRebaseMerge))
		buf.WriteString(fmt.Sprintf("\ttopics\t%v\n", c.After.Topics))
	}
	return buf.String()
}

func GetRepoChange(st *state.Repo, sp *spec.Repo) *RepoChange {
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
	return &RepoChange{
		Action: Update,
		Before: st,
		After:  after,
	}
}
