package ghrc

import (
	"bytes"
	"context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

var ghc *github.Client

func init() {
	accessToken := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	// initialize ghc var
	ghc = github.NewClient(tc)
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
	Topics      []Topic `yaml:"topics"`
	Labels      []Label `yaml:"labels"`
}
type Topic string

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
