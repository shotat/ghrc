package change

import (
	"bytes"
	"context"
	"fmt"

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
	buf.WriteString(fmt.Sprintf("%s Repo\n", string(c.Action)))
	diff := cmp.Diff(c.Before, c.After)
	buf.WriteString(diff)
	return buf.String()
}

func GetRepoChange(st *state.Repo, sp *spec.Repo) *RepoChange {
	after := &state.Repo{
		Name:             st.Name,
		Owner:            st.Owner,
		Description:      sp.Description,
		Homepage:         sp.Homepage,
		Private:          sp.Private,
		AllowSquashMerge: sp.AllowSquashMerge,
		AllowMergeCommit: sp.AllowMergeCommit,
		AllowRebaseMerge: sp.AllowRebaseMerge,
		Topics:           sp.Topics,
	}

	return &RepoChange{
		Action: Update,
		Before: st,
		After:  after,
	}
}
