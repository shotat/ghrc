package change

import (
	"github.com/shotat/ghrc/spec"
	"github.com/shotat/ghrc/state"
)

func GetLabelsChangeSet(st []state.Label, sp spec.Labels) []*LabelChange {
	if sp == nil {
		return nil
	}
	changes := make([]*LabelChange, 0)
	for _, spl := range sp {
		func(spl spec.Label) {
			for _, stl := range st {
				if stl.Name == spl.Name {
					// update existing label
					after := state.Label{

						Name:        spl.Name,
						Color:       spl.Color,
						Description: stl.Description,
					}
					if spl.Description != nil {
						after.Description = spl.Description
					}
					changes = append(changes, &LabelChange{
						Action: Update,
						Before: &stl,
						After:  &after,
					})
					return
				}
			}
			// new label
			changes = append(changes, &LabelChange{
				Action: Create,
				Before: nil,
				After: &state.Label{
					Name:        spl.Name,
					Color:       spl.Color,
					Description: spl.Description,
				},
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
