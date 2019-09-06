package status

import (
	"context"
	"github.com/google/go-github/v28/github"
)

type Label struct {
	repoOwner   string
	repoName    string
	ID          *int64
	Name        string
	Description *string
	Color       string
}

func (l *Label) Create(ctx context.Context) error {
	ghl := &github.Label{
		Name:        &l.Name,
		Color:       &l.Color,
		Description: l.Description,
	}
	_, _, err := ghc.Issues.CreateLabel(ctx, l.repoOwner, l.repoName, ghl)
	return err
}

func (l *Label) Change(ctx context.Context) error {
	ghl := &github.Label{
		Name:        &l.Name,
		Color:       &l.Color,
		Description: l.Description,
	}
	_, _, err := ghc.Issues.EditLabel(ctx, l.repoOwner, l.repoName, l.Name, ghl)
	return err
}

func (l *Label) Destroy(ctx context.Context) error {
	_, err := ghc.Issues.DeleteLabel(ctx, l.repoOwner, l.repoName, l.Name)
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
			ID:          label.ID,
			Name:        label.GetName(),
			Description: label.Description,
			Color:       label.GetColor(),
		}
	}
	return labels, nil
}
