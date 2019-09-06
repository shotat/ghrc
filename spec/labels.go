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

func LoadLabelsSpecFromSpec(states []state.Label) Labels {
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
