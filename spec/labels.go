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
		description := label.Description
		color := label.Color
		specs[i] = Label{
			Name:        label.Name,
			Description: &description,
			Color:       &color,
		}
	}
	return specs
}
