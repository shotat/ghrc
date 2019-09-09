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
		buf.WriteString(fmt.Sprintf("\tdescription\t%v\n", c.After.Description))
		buf.WriteString(fmt.Sprintf("\thomepage\t%v\n", c.After.Homepage))
		buf.WriteString(fmt.Sprintf("\tprivate\t%v\n", c.After.Private))
		buf.WriteString(fmt.Sprintf("\tallowSquashMerge\t%v\n", c.After.AllowSquashMerge))
		buf.WriteString(fmt.Sprintf("\tallowMergeCommit\t%v\n", c.After.AllowMergeCommit))
		buf.WriteString(fmt.Sprintf("\tallowRebaseMerge\t%v\n", c.After.AllowRebaseMerge))
		buf.WriteString(fmt.Sprintf("\ttopics\t%v\n", c.After.Topics))
	}
	return buf.String()
}

func GetRepoChange(st *state.Repo, sp *spec.Repo) *RepoChange {
	after := &state.Repo{
		ID:               st.ID,
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
