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

type RepositoryConfig struct {
	Metadata *metadata.Metadata `yaml:"metadata"`
	Spec     *spec.Spec         `yaml:"spec"`
}

func (c *RepositoryConfig) ToYAML() (string, error) {
	buf := bytes.NewBuffer(nil)
	err := yaml.NewEncoder(buf).Encode(c)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func LoadFromFile(path string) (*RepositoryConfig, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	conf := new(RepositoryConfig)
	if err := yaml.NewDecoder(bytes.NewReader(buf)).Decode(conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func Import(ctx context.Context, owner string, name string) (*RepositoryConfig, error) {
	conf := new(RepositoryConfig)
	meta := &metadata.Metadata{
		Owner: owner,
		Name:  name,
	}

	conf.Metadata = meta

	repo, err := state.FindRepo(owner, name)
	if err != nil {
		return nil, err
	}
	sp := new(spec.Spec)
	sp.Repo.Homepage = repo.Homepage
	sp.Repo.Description = repo.Description
	sp.Repo.Private = repo.Private
	sp.Repo.Topics = repo.Topics
	sp.Repo.AllowSquashMerge = repo.AllowSquashMerge
	sp.Repo.AllowMergeCommit = repo.AllowMergeCommit
	sp.Repo.AllowRebaseMerge = repo.AllowRebaseMerge

	labels, err := state.FindLabels(ctx, owner, name)
	if err != nil {
		return nil, err
	}
	if labels != nil {
		sp.Labels = make([]spec.Label, len(labels))
		for i, label := range labels {
			sp.Labels[i] = spec.Label{
				Name:        label.Name,
				Description: label.Description,
				Color:       label.Color,
			}
		}
	}

	// TODO
	// spec.Protections = repo.Protections

	conf.Spec = sp

	return conf, nil
}

func (rc *RepositoryConfig) calculateChangeSet(ctx context.Context) (change.ChangeSet, error) {
	repo, err := state.FindRepo(rc.Metadata.Owner, rc.Metadata.Name)
	if err != nil {
		return nil, err
	}
	labels, err := state.FindLabels(ctx, rc.Metadata.Owner, rc.Metadata.Name)
	if err != nil {
		return nil, err
	}

	changeSet := make(change.ChangeSet, 0)
	changeSet = append(changeSet, rc.Spec.Repo.GetRepoChange(repo))
	for _, labelChange := range rc.Spec.Labels.GetLabelsChangeSet(labels) {
		changeSet = append(changeSet, labelChange)
	}
	return changeSet, nil
}

func (rc *RepositoryConfig) Plan(ctx context.Context) error {
	cs, err := rc.calculateChangeSet(ctx)
	if err != nil {
		return err
	}

	for _, c := range cs {
		fmt.Println(c)
	}
	return nil
}

func (rc *RepositoryConfig) Apply(ctx context.Context) error {
	cs, err := rc.calculateChangeSet(ctx)
	if err != nil {
		return err
	}

	for _, c := range cs {
		if err := c.Apply(ctx, rc.Metadata.Owner, rc.Metadata.Name); err != nil {
			return err
		}
	}
	return nil
}
