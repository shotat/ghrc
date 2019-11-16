package change

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/google/go-cmp/cmp"
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
	buf.WriteString(c.subject())
	diff := cmp.Diff(c.Before, c.After)
	buf.WriteString(diff)
	return buf.String()
}

func (c *LabelChange) subject() string {
	var name string
	switch c.Action {
	case Create, Update:
		name = c.After.Name
	case Delete:
		name = c.Before.Name
	}
	return fmt.Sprintf("%s Label: %s\n", string(c.Action), name)
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

func getLabelChange(st *state.Label, sp *spec.Label) *LabelChange {
	if st == nil {
		after := &state.Label{
			Name:        sp.Name,
			Description: "",
			Color:       "f0f0f0", // TODO: auto generation
		}
		if sp.Description != nil {
			after.Description = *sp.Description
		}
		if sp.Color != nil {
			after.Color = *sp.Color
		}
		// New
		return &LabelChange{
			Action: Create,
			Before: nil,
			After:  after,
		}
	}

	if sp == nil {
		// Delete
		return &LabelChange{
			Action: Delete,
			Before: st,
			After:  nil,
		}
	}

	// Update
	after := &state.Label{
		Name:        st.Name,
		Description: st.Description,
		Color:       st.Color,
	}
	if sp.Description != nil {
		after.Description = *sp.Description
	}

	if sp.Color != nil {
		after.Color = *sp.Color
	}

	return &LabelChange{
		Action: Update,
		Before: st,
		After:  after,
	}
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
					ch := getLabelChange(&stl, &spl)
					changes = append(changes, ch)
					return
				}
			}
			// new label
			ch := getLabelChange(nil, &spl)
			changes = append(changes, ch)
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
			ch := getLabelChange(&stl, nil)
			changes = append(changes, ch)
		}(stl)
	}
	return changes
}
