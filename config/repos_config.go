package config

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/shotat/ghrc/metadata"
	"github.com/shotat/ghrc/spec"
	"github.com/shotat/ghrc/status"
	yaml "gopkg.in/yaml.v2"
)

type RepositoryConfig struct {
	Metadata *metadata.RepositoryMetadata `yaml:"metadata"`
	Spec     *spec.RepositorySpec         `yaml:"spec"`
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
	repo, err := status.FindRepositoryStatus(owner, name)
	if err != nil {
		return nil, err
	}
	conf := new(RepositoryConfig)
	meta := &metadata.RepositoryMetadata{
		Owner: owner,
		Name:  name,
	}

	conf.Metadata = meta

	rs := new(spec.RepositorySpec)
	rs.Homepage = repo.Homepage
	rs.Description = repo.Description
	rs.Private = repo.Private
	rs.Topics = repo.Topics
	rs.AllowSquashMerge = repo.AllowSquashMerge
	rs.AllowMergeCommit = repo.AllowMergeCommit
	rs.AllowRebaseMerge = repo.AllowRebaseMerge

	if repo.Labels != nil {
		rs.Labels = make([]spec.Label, len(repo.Labels))
		for i, label := range repo.Labels {
			rs.Labels[i] = spec.Label{
				Name:        label.Name,
				Description: label.Description,
				Color:       label.Color,
			}
		}
	}

	// TODO
	// spec.Protections = repo.Protections

	conf.Spec = rs

	return conf, nil
}

// TODO
func (rc *RepositoryConfig) Plan(ctx context.Context) error {
	return nil
}

func (rc *RepositoryConfig) Apply(ctx context.Context) error {
	repo, err := status.FindRepositoryStatus(rc.Metadata.Owner, rc.Metadata.Name)
	if err != nil {
		return err
	}

	rc.Spec.Patch(repo)

	return repo.Apply(ctx)
}
