package change

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/shotat/ghrc/spec"
	"github.com/shotat/ghrc/state"
)

type LabelChange struct {
	Action Action
	Before *state.Label
	After  *state.Label
}

func (c *LabelChange) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("%s Label\n", string(c.Action)))
	switch c.Action {
	case Create:
		buf.WriteString(fmt.Sprintf("\tname\t%v\n", c.After.Name))
		buf.WriteString(fmt.Sprintf("\tcolor\t%v\n", c.After.Color))
	case Update:
		buf.WriteString(fmt.Sprintf("\tname\t%v\n", c.After.Name))
		buf.WriteString(fmt.Sprintf("\tcolor\t%v\n", c.After.Color))
	case Delete:
		buf.WriteString(fmt.Sprintf("\tname\t%v\n", c.Before.Name))
		buf.WriteString(fmt.Sprintf("\tcolor\t%v\n", c.Before.Color))
	}
	return buf.String()
}

func (c *LabelChange) Apply(ctx context.Context, repoOwner string, repoName string) error {
	err := errors.New("unexpected error")
	switch c.Action {
	case Create:
		err = c.After.Create(ctx, repoOwner, repoName)
	case Delete:
		err = c.Before.Destroy(ctx, repoOwner, repoName)
	case Update:
		err = c.After.Update(ctx, repoOwner, repoName)
	}
	return err
}

func GetLabelChangeSet(st []state.Label, sp spec.Labels) []*LabelChange {
	changes := make([]*LabelChange, 0)
	if sp == nil {
		return changes
	}
	for _, spl := range sp {
		func(spl spec.Label) {
			for _, stl := range st {
				if stl.Name == spl.Name {
					// update existing label
					after := spl.ToState(&stl)
					changes = append(changes, &LabelChange{
						Action: Update,
						Before: &stl,
						After:  after,
					})
					return
				}
			}
			// new label
			changes = append(changes, &LabelChange{
				Action: Create,
				Before: nil,
				After:  spl.ToState(nil),
			})
			return
		}(spl)
	}
	for _, stl := range st {
		func(stl state.Label) {
			for _, spl := range sp {
				if stl.Name == spl.Name {
					return
				}
			}

			// deletion
			changes = append(changes, &LabelChange{
				Action: Delete,
				Before: &stl,
				After:  nil,
			})
		}(stl)
	}
	return changes
}
