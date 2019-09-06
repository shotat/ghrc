package status

import (
	"context"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v28/github"
)

type Repo struct {
	ID               int64
	Name             string
	Owner            string
	Description      *string
	Homepage         *string
	Private          *bool
	AllowSquashMerge *bool
	AllowMergeCommit *bool
	AllowRebaseMerge *bool

	Topics []string
}

func FindRepo(owner string, name string) (*Repo, error) {
	ctx := context.Background()
	repo, _, err := ghc.Repositories.Get(ctx, owner, name)
	if err != nil {
		return nil, err
	}
	return &Repo{
		ID:               repo.GetID(),
		Name:             repo.GetName(),
		Owner:            repo.GetOwner().GetLogin(),
		Homepage:         repo.Homepage,
		Description:      repo.Description,
		Private:          repo.Private,
		Topics:           repo.Topics,
		AllowSquashMerge: repo.AllowSquashMerge,
		AllowMergeCommit: repo.AllowMergeCommit,
		AllowRebaseMerge: repo.AllowRebaseMerge,
	}, nil
}

func (s *Repo) Diff(t *Repo) string {
	return cmp.Diff(s, t)
}

func (s *Repo) Update(ctx context.Context) error {
	repo := new(github.Repository)

	repo.Name = &s.Name
	repo.Description = s.Description
	repo.Homepage = s.Description
	repo.Private = s.Private
	repo.AllowRebaseMerge = s.AllowRebaseMerge
	repo.AllowSquashMerge = s.AllowSquashMerge
	repo.AllowMergeCommit = s.AllowMergeCommit

	_, _, err := ghc.Repositories.Edit(ctx, s.Owner, s.Name, repo)
	if err != nil {
		return err
	}

	// TODO topic

	return nil
}
