package change

import (
	"context"

	"github.com/shotat/ghrc/spec"
	"github.com/shotat/ghrc/state"
)

type ProtectionChange struct {
	Action Action
	Before *state.Protection
	After  *state.Protection
}

func (c *ProtectionChange) String() string {
	return "TODO"
}

func (c *ProtectionChange) Apply(ctx context.Context, repoOwner string, repoName string) error {
	return nil
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
					after := state.Protection{
						// TODO
					}
					changes = append(changes, &ProtectionChange{
						Action: Update,
						Before: &stp,
						After:  &after,
					})
					return
				}
			}

			// new protection
			changes = append(changes, &ProtectionChange{
				Action: Create,
				Before: nil,
				After:  &state.Protection{
					// TODO
				},
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
