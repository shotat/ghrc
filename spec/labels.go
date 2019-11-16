package spec

import (
	"github.com/shotat/ghrc/state"
)

type Label struct {
	Name        string  `yaml:"name"`
	Description *string `yaml:"description"`
	Color       *string `yaml:"color"` // If omitted, random color is generated
}

type Labels []Label

func LoadLabelsSpecFromState(states []state.Label) Labels {
	specs := make([]Label, len(states))
	for i, label := range states {
		specs[i] = Label{
			Name:        label.Name,
			Description: &label.Description,
			Color:       &label.Color,
		}
	}
	return specs
}

/*
// ToState generates a new state
func (sp *Label) ToState() *state.Label {
	newState := &state.Label{}
	newState.Name = sp.Name
	newState.Color = sp.Color
	newState.Description = sp.Description
	return newState
}
*/
