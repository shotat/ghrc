package ghrc

import (
	"context"
	"github.com/google/go-github/v28/github"
)

type Protection struct {
	Branch                     *string                     `yaml:"branch"`
	RequiredStatusCheck        *RequiredStatusCheck        `yaml:"requiredStatusCheck"`
	EnforceAdmins              *bool                       `yaml:"enforceAdmins"`
	RequiredPullRequestReviews *RequiredPullRequestReviews `yaml:"requiredPullRequestReviews"`
	Restrictions               *Restrictions               `yaml:"restrictions"`
}

type RequiredPullRequestReviews struct {
	DismissalRestrictions        *Restrictions `yaml:"dismissalRestrictions"`
	DismissStaleReviews          bool          `yaml:"dismissStaleReviews"`
	RequireCodeOwnerReviews      bool          `yaml:"requireCodeOwnerReviews"`
	RequiredApprovingReviewCount int           `yaml:"requiredApprovingReviewCount"`
}

type Restrictions struct {
	Users []string `yaml:"users"`
	Teams []string `yaml:"teams"`
}

type RequiredStatusCheck struct {
	Strict   bool     `yaml:"strict"`
	Contexts []string `yaml:"contexts"`
}

func findProtections(meta *RepositoryMetadata) ([]Protection, error) {
	ctx := context.Background()
	protected := true
	opt := &github.BranchListOptions{
		Protected: &protected,
	}
	protectedBranches, _, err := ghc.Repositories.ListBranches(ctx, meta.Owner, meta.Name, opt)
	if err != nil {
		return nil, err
	}

	protections := make([]Protection, len(protectedBranches))
	for i, pb := range protectedBranches {
		p, _, err := ghc.Repositories.GetBranchProtection(ctx, meta.Owner, meta.Name, pb.GetName())
		if err != nil {
			return nil, err
		}
		protections[i] = Protection{
			Branch: pb.Name,

			RequiredStatusCheck: &RequiredStatusCheck{
				Strict:   p.GetRequiredStatusChecks().Strict,
				Contexts: p.GetRequiredStatusChecks().Contexts,
			},
			EnforceAdmins: &p.GetEnforceAdmins().Enabled,
			RequiredPullRequestReviews: &RequiredPullRequestReviews{
				DismissalRestrictions:        nil, // TODO
				DismissStaleReviews:          p.GetRequiredPullRequestReviews().DismissStaleReviews,
				RequireCodeOwnerReviews:      p.GetRequiredPullRequestReviews().RequireCodeOwnerReviews,
				RequiredApprovingReviewCount: p.GetRequiredPullRequestReviews().RequiredApprovingReviewCount,
			},
			Restrictions: nil, // TODO
		}
	}
	return protections, err
}
