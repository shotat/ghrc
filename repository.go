package ghrc

import (
	"bytes"
	"context"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"github.com/google/go-github/v28/github"
)

var ghc *github.Client

func init() {
	client := github.NewClient(nil)
	ghc = client
}

func FindRepository(meta *RepositoryMetadata) (*github.Repository, error) {
	ctx := context.Background()
	repo, _, err := ghc.Repositories.Get(ctx, meta.Owner, meta.Name)
	if err != nil {
		return nil, err
	}
	return repo, nil
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
	Description *string `yaml:"description"`
	Private     *bool   `yaml:"private"`
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

func (rc *RepositoryConfig) Apply() error {
	ctx := context.Background()
	repo, err := FindRepository(rc.Metadata)
	if err != nil {
		return err
	}
	rc.Spec.Patch(repo)
	_, _, err = ghc.Repositories.Edit(ctx, rc.Metadata.Owner, rc.Metadata.Name, repo)
	return err
}
