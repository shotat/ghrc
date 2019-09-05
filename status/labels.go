package status

import (
	"context"

	"github.com/google/go-github/v28/github"
)

type Label struct {
	ID          int64
	Name        string
	Description *string
	Color       string
}

func findLabels(ctx context.Context, owner string, repo string) ([]*github.Label, error) {
	labels, _, err := ghc.Issues.ListLabels(ctx, owner, repo, nil)
	return labels, err
}
