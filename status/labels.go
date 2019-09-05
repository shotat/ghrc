package status

import (
	"context"

	"github.com/google/go-github/v28/github"
)

func findLabels(owner string, repo string) ([]*github.Label, error) {
	ctx := context.Background()
	labels, _, err := ghc.Issues.ListLabels(ctx, owner, repo, nil)
	return labels, err
}
