package spec

import (
	"github.com/shotat/ghrc/change"
	"github.com/shotat/ghrc/status"
)

type Label struct {
	Name        string  `yaml:"name"`
	Description *string `yaml:"description,omitempty"`
	Color       string  `yaml:"color"`
}

type Labels []Label

func (sp Labels) GetLabelsChangeSet(st []status.Label) []change.LabelChange {
	if sp == nil {
		return nil
	}
	changes := make([]change.LabelChange, 0)
	for _, spl := range sp {
		func() {
			for _, stl := range st {
				if stl.Name == spl.Name {
					// update existing label
					after := status.Label{
						Name:        spl.Name,
						Color:       spl.Color,
						Description: stl.Description,
					}
					if spl.Description != nil {
						after.Description = spl.Description
					}
					changes = append(changes, change.LabelChange{
						Action: change.Update,
						Before: &stl,
						After:  &after,
					})
					return
				}
			}
			// new label
			changes = append(changes, change.LabelChange{
				Action: change.Create,
				Before: nil,
				After: &status.Label{
					Name:        spl.Name,
					Color:       spl.Color,
					Description: spl.Description,
				},
			})
			return
		}()
	}
	for _, stl := range st {
		func() {
			for _, spl := range sp {
				if stl.Name == spl.Name {
					return
				}
			}

			// deletion
			changes = append(changes, change.LabelChange{
				Action: change.Delete,
				Before: &stl,
				After:  nil,
			})
			return
		}()
	}
	return changes
}
