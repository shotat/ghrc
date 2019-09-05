package status

import (
	"context"
	"github.com/google/go-github/v28/github"
)

type RepositoryStatus struct {
	ID               *int
	Name             *string
	Owner            *string
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
	status := new(RepositoryStatus)
	status.Homepage = repo.Homepage
	status.Description = repo.Description
	status.Private = repo.Private
	status.Topics = repo.Topics
	status.AllowSquashMerge = repo.AllowSquashMerge
	status.AllowMergeCommit = repo.AllowMergeCommit
	status.AllowRebaseMerge = repo.AllowRebaseMerge

	labels, err := findLabels(owner, name)
	if err != nil {
		return nil, err
	}
	for _, label := range labels {
		status.Labels = append(status.Labels, Label{
			Name:        label.Name,
			Description: label.Description,
			Color:       label.Color,
		})
	}

	protections, err := findProtections(owner, name)
	if err != nil {
		return nil, err
	}
	status.Protections = protections

	return status, nil
}

func (s *RepositoryStatus) Sync(ctx context.Context) error {
	repo := new(github.Repository)
	_, _, err := ghc.Repositories.Edit(ctx, *s.Owner, *s.Name, repo)
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

type Label struct {
	ID          *int
	Name        *string
	Description *string
	Color       *string
}
