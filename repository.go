package ghrc

import (
	"bytes"
	"context"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/google/go-github/v28/github"
)

func FindRepository(meta *RepositoryMetadata) (*github.Repository, error) {
	ctx := context.Background()
	repo, _, err := ghc.Repositories.Get(ctx, meta.Owner, meta.Name)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func ExportConfig(meta *RepositoryMetadata) (*RepositoryConfig, error) {
	ctx := context.Background()
	repo, _, err := ghc.Repositories.Get(ctx, meta.Owner, meta.Name)
	if err != nil {
		return nil, err
	}
	labels, _, err := ghc.Issues.ListLabels(ctx, meta.Owner, meta.Name, nil)
	if err != nil {
		return nil, err
	}
	branches, _, err := ghc.Repositories.ListBranches(ctx, meta.Owner, meta.Name, nil)
	var protectedBranches []*github.Branch
	for _, b := range branches {
		if b.GetProtected() {
			protectedBranches = append(protectedBranches, b)
		}
	}

	conf := new(RepositoryConfig)
	conf.Metadata = meta
	// Spec
	spec := new(RepositorySpec)
	spec.Description = repo.Description
	spec.Private = repo.Private
	spec.Topics = repo.Topics
	for _, label := range labels {
		spec.Labels = append(spec.Labels, Label{
			Name:        label.Name,
			Description: label.Description,
			Color:       label.Color,
		})
	}
	for _, pb := range protectedBranches {
		spec.Protections = append(spec.Protections, Protection{
			Branch: pb.Name,
		})
	}

	conf.Spec = spec
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
	Description *string      `yaml:"description"`
	Private     *bool        `yaml:"private"`
	Topics      []string     `yaml:"topics"`
	Labels      []Label      `yaml:"labels"`
	Protections []Protection `yaml:"protections"`
}

type Protection struct {
	Branch        *string `yaml:"branch"`
	EnforceAdmins *bool   `yaml:"enforceAdmins"`
}

type Label struct {
	Name        *string `yaml:"name"`
	Description *string `yaml:"description"`
	Color       *string `yaml:"color"`
}

func (rs *RepositorySpec) Patch(repo *github.Repository) {
	repo.Description = rs.Description
	repo.Private = rs.Private
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
func (rc *RepositoryConfig) Plan() error {
	return nil
}

func (rc *RepositoryConfig) Apply() error {
	ctx := context.Background()
	repo, err := FindRepository(rc.Metadata)
	if err != nil {
		return err
	}
	fmt.Println(rc.Spec.Topics)
	_, _, err = ghc.Repositories.ReplaceAllTopics(ctx, rc.Metadata.Owner, rc.Metadata.Name, rc.Spec.Topics)
	if err != nil {
		return err
	}

	rc.Spec.Patch(repo)
	_, _, err = ghc.Repositories.Edit(ctx, rc.Metadata.Owner, rc.Metadata.Name, repo)
	return err
}
