package status

import (
	"context"
)

type Label struct {
	ID          int64
	Name        string
	Description *string
	Color       string
}

func findLabels(ctx context.Context, owner string, repo string) ([]Label, error) {
	ghLabels, _, err := ghc.Issues.ListLabels(ctx, owner, repo, nil)
	if err != nil {
		return nil, err
	}
	labels := make([]Label, len(ghLabels))
	for i, label := range ghLabels {
		labels[i] = Label{
			ID:          label.GetID(),
			Name:        label.GetName(),
			Description: label.Description,
			Color:       label.GetColor(),
		}
	}
	return labels, nil
}
