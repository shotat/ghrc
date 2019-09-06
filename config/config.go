package config

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/shotat/ghrc/metadata"
	"github.com/shotat/ghrc/spec"
	"github.com/shotat/ghrc/state"
	yaml "gopkg.in/yaml.v2"
)

// Config represents a desired remote configuration.
type Config struct {
	Metadata *metadata.Metadata `yaml:"metadata"`
	Spec     *spec.Spec         `yaml:"spec"`
}

// ToYAML serialize Config to the YAML format.
func (c *Config) ToYAML() (string, error) {
	buf := bytes.NewBuffer(nil)
	err := yaml.NewEncoder(buf).Encode(c)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// LoadFromFile loads a config file from `path`
// the format must be YAML.
func LoadFromFile(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	conf := new(Config)
	if err := yaml.NewDecoder(bytes.NewReader(buf)).Decode(conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// Import remote state to config
func Import(ctx context.Context, owner string, name string) (*Config, error) {
	conf := new(Config)
	conf.Metadata = metadata.NewMetadata(owner, name)
	conf.Spec = new(spec.Spec)

	repo, err := state.FindRepo(ctx, owner, name)
	if err != nil {
		return nil, err
	}

	labels, err := state.FindLabels(ctx, owner, name)
	if err != nil {
		return nil, err
	}

	conf.Spec.Repo = spec.LoadRepoSpecFromState(repo)
	conf.Spec.Labels = spec.LoadLabelsSpecFromSpec(labels)

	// TODO
	// spec.Protections = repo.Protections

	return conf, nil
}
