package ghrc

import (
	"context"

	"github.com/google/go-github/v28/github"
)

func findLabels(meta *RepositoryMetadata) ([]*github.Label, error) {
	ctx := context.Background()
	labels, _, err := ghc.Issues.ListLabels(ctx, meta.Owner, meta.Name, nil)
	return labels, err
}
