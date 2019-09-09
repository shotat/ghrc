package spec

import (
	"github.com/shotat/ghrc/state"
)

type Label struct {
	Name        string  `yaml:"name"`
	Description *string `yaml:"description,omitempty"`
	Color       string  `yaml:"color"`
}

type Labels []Label

func LoadLabelsSpecFromState(states []state.Label) Labels {
	specs := make([]Label, len(states))
	for i, label := range states {
		specs[i] = Label{
			Name:        label.Name,
			Description: label.Description,
			Color:       label.Color,
		}
	}
	return specs
}

// ToState merge state and spec to generate new state
func (sp *Label) ToState(base *state.Label) *state.Label {
	newState := &state.Label{}
	// initialize
	if base != nil {
		newState.Description = base.Description
	}
	newState.Name = sp.Name
	newState.Color = sp.Color
	return newState
}
