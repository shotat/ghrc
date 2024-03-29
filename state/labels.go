package state

import (
	"context"
	"github.com/google/go-github/v28/github"
)

type Label struct {
	Name        string
	Description string
	Color       string
}

func (l *Label) Create(ctx context.Context, repoOwner string, repoName string) error {
	ghl := &github.Label{
		Name:        &l.Name,
		Color:       &l.Color,
		Description: &l.Description,
	}
	_, _, err := ghc.Issues.CreateLabel(ctx, repoOwner, repoName, ghl)
	return err
}

func (l *Label) Update(ctx context.Context, repoOwner string, repoName string) error {
	ghl := &github.Label{
		Name:        &l.Name,
		Color:       &l.Color,
		Description: &l.Description,
	}
	_, _, err := ghc.Issues.EditLabel(ctx, repoOwner, repoName, l.Name, ghl)
	return err
}

func (l *Label) Destroy(ctx context.Context, repoOwner string, repoName string) error {
	_, err := ghc.Issues.DeleteLabel(ctx, repoOwner, repoName, l.Name)
	return err
}

func FindLabels(ctx context.Context, owner string, repo string) ([]Label, error) {
	ghLabels, _, err := ghc.Issues.ListLabels(ctx, owner, repo, nil)
	if err != nil {
		return nil, err
	}
	labels := make([]Label, len(ghLabels))
	for i, label := range ghLabels {
		labels[i] = Label{
			Name:        label.GetName(),
			Description: label.GetDescription(),
			Color:       label.GetColor(),
		}
	}
	return labels, nil
}
