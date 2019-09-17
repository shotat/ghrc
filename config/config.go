package config

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/shotat/ghrc/change"
	"github.com/shotat/ghrc/metadata"
	"github.com/shotat/ghrc/spec"
	"github.com/shotat/ghrc/state"
	yaml "gopkg.in/yaml.v2"
)

// Config represents a desired remote configuration.
type Config struct {
	Metadata metadata.Metadata `yaml:"metadata"`
	Spec     spec.Spec         `yaml:"spec"`
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
	conf.Metadata = *metadata.NewMetadata(owner, name)
	conf.Spec = spec.Spec{}

	repo, err := state.FindRepo(ctx, owner, name)
	if err != nil {
		return nil, err
	}

	labels, err := state.FindLabels(ctx, owner, name)
	if err != nil {
		return nil, err
	}

	protections, err := state.FindProtections(ctx, owner, name)
	if err != nil {
		return nil, err
	}

	conf.Spec.Repo = spec.LoadRepoSpecFromState(repo)
	conf.Spec.Labels = spec.LoadLabelsSpecFromState(labels)
	conf.Spec.Protections = spec.LoadProtectionsSpecFromState(protections)

	return conf, nil
}

func (c *Config) CalculateChangeSet(ctx context.Context) (change.ChangeSet, error) {
	repo, err := state.FindRepo(ctx, c.Metadata.Owner, c.Metadata.Name)
	if err != nil {
		return nil, err
	}
	labels, err := state.FindLabels(ctx, c.Metadata.Owner, c.Metadata.Name)
	if err != nil {
		return nil, err
	}
	protections, err := state.FindProtections(ctx, c.Metadata.Owner, c.Metadata.Name)
	if err != nil {
		return nil, err
	}

	changeSet := make(change.ChangeSet, 0)
	if ch := change.GetRepoChange(repo, c.Spec.Repo); ch != nil {
		changeSet = append(changeSet, ch)
	}
	for _, labelChange := range change.GetLabelChangeSet(labels, c.Spec.Labels) {
		changeSet = append(changeSet, labelChange)
	}
	for _, protectionChange := range change.GetProtectionChangeSet(protections, c.Spec.Protections) {
		changeSet = append(changeSet, protectionChange)
	}
	return changeSet, nil
}

// Plan shows the expected changes without changing actual states.
func (c *Config) Plan(ctx context.Context) error {
	cs, err := c.CalculateChangeSet(ctx)
	if err != nil {
		return err
	}

	for _, ch := range cs {
		fmt.Println(ch)
	}
	return nil
}

// Apply changes the remote configurations based on this Config.
func (c *Config) Apply(ctx context.Context) error {
	cs, err := c.CalculateChangeSet(ctx)
	if err != nil {
		return err
	}

	for _, ch := range cs {
		if err := ch.Apply(ctx, c.Metadata.Owner, c.Metadata.Name); err != nil {
			return err
		}
	}
	return nil
}
