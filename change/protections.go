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

type ProtectionChange struct {
	Action Action
	Before *state.Protection
	After  *state.Protection
}

// FIXME: duplicated
func (c *ProtectionChange) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(c.subject())
	diff := cmp.Diff(c.Before, c.After)
	buf.WriteString(diff)
	return buf.String()
}

func (c *ProtectionChange) subject() string {
	var name string
	switch c.Action {
	case Create, Update:
		name = c.After.Branch
	case Delete:
		name = c.Before.Branch
	}
	return fmt.Sprintf("%s Protection: %s\n", string(c.Action), name)
}

func (c *ProtectionChange) Apply(ctx context.Context, repoOwner string, repoName string) error {
	err := errors.New("unexpected error")
	switch c.Action {
	case Create, Update:
		err = c.After.Update(ctx, repoOwner, repoName)
	case Delete:
		err = c.Before.Destroy(ctx, repoOwner, repoName)
	}
	return err
}

func GetProtectionChangeSet(st []state.Protection, sp spec.Protections) []*ProtectionChange {
	changes := make([]*ProtectionChange, 0)
	if sp == nil {
		return changes
	}

	for _, spp := range sp {
		func(spp spec.Protection) {
			for _, stp := range st {
				if stp.Branch == spp.Branch {
					// update
					changes = append(changes, &ProtectionChange{
						Action: Update,
						Before: &stp,
						After:  spp.ToState(),
					})
					return
				}
			}

			// new protection
			changes = append(changes, &ProtectionChange{
				Action: Create,
				Before: nil,
				After:  spp.ToState(),
			})
		}(spp)
	}

	for _, stp := range st {
		func(stp state.Protection) {
			for _, spp := range sp {
				if stp.Branch == spp.Branch {
					return
				}
			}

			// deletion
			changes = append(changes, &ProtectionChange{
				Action: Delete,
				Before: &stp,
				After:  nil,
			})
		}(stp)
	}
	return changes
}
