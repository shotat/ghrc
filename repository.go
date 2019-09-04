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

func ExportConfig(repo *github.Repository) (*RepositoryConfig, error) {
	conf := new(RepositoryConfig)
	spec := new(RepositorySpec)
	spec.Description = repo.Description
	spec.Private = repo.Private
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
	Description *string  `yaml:"description"`
	Private     *bool    `yaml:"private"`
	Topics      []string `yaml:"topics"`
	Labels      []Label  `yaml:"labels"`
}

type Label struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Color       string `yaml:"color"`
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
