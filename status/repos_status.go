package status

import (
	"context"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v28/github"
)

type RepositoryStatus struct {
	ID               int64
	Name             string
	Owner            string
	Description      *string
	Homepage         *string
	Private          *bool
	AllowSquashMerge *bool
	AllowMergeCommit *bool
	AllowRebaseMerge *bool

	Topics      []string
	Labels      []Label
	Protections []Protection
}

func FindRepositoryStatus(owner string, name string) (*RepositoryStatus, error) {
	ctx := context.Background()
	repo, _, err := ghc.Repositories.Get(ctx, owner, name)
	if err != nil {
		return nil, err
	}
	// Spec
	status := &RepositoryStatus{
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
	}

	labels, err := findLabels(ctx, owner, name)
	if err != nil {
		return nil, err
	}
	status.Labels = labels

	protections, err := findProtections(owner, name)
	if err != nil {
		return nil, err
	}
	status.Protections = protections

	return status, nil
}

func (s *RepositoryStatus) Diff(t *RepositoryStatus) string {
	return cmp.Diff(s, t)
}

func (s *RepositoryStatus) Apply(ctx context.Context) error {
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

	/*
		if rc.Spec.Topics != nil {
			_, _, err = ghc.Repositories.ReplaceAllTopics(ctx, rc.Metadata.Owner, rc.Metadata.Name, rc.Spec.Topics)
			if err != nil {
				return err
			}
		}
	*/

	// TODO label

	// TODO protections

	return nil
}
