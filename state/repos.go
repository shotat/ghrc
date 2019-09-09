package state

import (
	"context"

	"github.com/google/go-github/v28/github"
)

type Repo struct {
	ID               int64
	Name             string // required?
	Owner            string
	Description      string
	Homepage         string
	Private          bool
	AllowSquashMerge bool
	AllowMergeCommit bool
	AllowRebaseMerge bool
	Topics           []string
}

func FindRepo(ctx context.Context, owner string, name string) (*Repo, error) {
	repo, _, err := ghc.Repositories.Get(ctx, owner, name)
	if err != nil {
		return nil, err
	}
	return &Repo{
		ID:               repo.GetID(),
		Name:             repo.GetName(),
		Owner:            repo.GetOwner().GetLogin(),
		Homepage:         repo.GetHomepage(),
		Description:      repo.GetDescription(),
		Private:          repo.GetPrivate(),
		Topics:           repo.Topics,
		AllowSquashMerge: repo.GetAllowSquashMerge(),
		AllowMergeCommit: repo.GetAllowMergeCommit(),
		AllowRebaseMerge: repo.GetAllowRebaseMerge(),
	}, nil
}

func (s *Repo) Update(ctx context.Context, repoOwner string, repoName string) error {
	repo := new(github.Repository)
	repo.Name = &s.Name
	repo.Description = &s.Description
	repo.Homepage = &s.Homepage
	repo.Private = &s.Private
	repo.AllowRebaseMerge = &s.AllowRebaseMerge
	repo.AllowSquashMerge = &s.AllowSquashMerge
	repo.AllowMergeCommit = &s.AllowMergeCommit

	if _, _, err := ghc.Repositories.Edit(ctx, repoOwner, repoName, repo); err != nil {
		return err
	}

	if s.Topics != nil {
		if _, _, err := ghc.Repositories.ReplaceAllTopics(ctx, s.Owner, s.Name, s.Topics); err != nil {
			return err
		}
	}

	return nil
}
