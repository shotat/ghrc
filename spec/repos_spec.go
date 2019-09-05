package spec

import (
	"bytes"
	"context"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/shotat/ghrc/status"
)

func ImportConfig(ctx context.Context, meta *RepositoryMetadata) (*RepositoryConfig, error) {
	repo, err := status.FindRepositoryStatus(meta.Owner, meta.Name)
	if err != nil {
		return nil, err
	}
	conf := new(RepositoryConfig)
	conf.Metadata = meta

	// Spec
	spec := new(RepositorySpec)
	spec.Homepage = repo.Homepage
	spec.Description = repo.Description
	spec.Private = repo.Private
	spec.Topics = repo.Topics
	spec.AllowSquashMerge = repo.AllowSquashMerge
	spec.AllowMergeCommit = repo.AllowMergeCommit
	spec.AllowRebaseMerge = repo.AllowRebaseMerge

	if repo.Labels != nil {
		spec.Labels = make([]Label, len(repo.Labels))
		for i, label := range repo.Labels {
			spec.Labels[i] = Label{
				Name:        label.Name,
				Description: label.Description,
				Color:       label.Color,
			}
		}
	}

	// TODO
	// spec.Protections = repo.Protections
	// conf.Spec = spec
	return conf, nil
}

type RepositoryConfig struct {
	Metadata *RepositoryMetadata `yaml:"metadata"`
	Spec     *RepositorySpec     `yaml:"spec"`
}

type RepositoryMetadata struct {
	Owner string `yaml:"owner"`
	Name  string `yaml:"name"`
}

type RepositorySpec struct {
	Description      *string `yaml:"description"`
	Homepage         *string `yaml:"homepage"`
	Private          *bool   `yaml:"private"`
	AllowSquashMerge *bool   `yaml:"allowSquashMerge"`
	AllowMergeCommit *bool   `yaml:"allowMergeCommit"`
	AllowRebaseMerge *bool   `yaml:"allowRebaseMerge"`

	Topics      []string     `yaml:"topics"`
	Labels      []Label      `yaml:"labels"`
	Protections []Protection `yaml:"protections"`
}

type Label struct {
	Name        *string `yaml:"name"`
	Description *string `yaml:"description"`
	Color       *string `yaml:"color"`
}

func (rs *RepositorySpec) Patch(repo *status.RepositoryStatus) {
	repo.Description = rs.Description
	repo.Private = rs.Private
	repo.Homepage = rs.Homepage
	repo.AllowSquashMerge = rs.AllowSquashMerge
	repo.AllowMergeCommit = rs.AllowMergeCommit
	repo.AllowRebaseMerge = rs.AllowRebaseMerge
}

func LoadRepositoryConfigFromFile(path string) (*RepositoryConfig, error) {
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

	return nil
}
