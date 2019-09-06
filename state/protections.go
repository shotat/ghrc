package state

import (
	"context"
	"github.com/google/go-github/v28/github"
)

type Protection struct {
	Branch                     *string
	RequiredStatusCheck        *RequiredStatusCheck
	EnforceAdmins              *bool
	RequiredPullRequestReviews *RequiredPullRequestReviews
	Restrictions               *Restrictions
}

type RequiredPullRequestReviews struct {
	DismissalRestrictions        *Restrictions
	DismissStaleReviews          bool
	RequireCodeOwnerReviews      bool
	RequiredApprovingReviewCount int
}

type Restrictions struct {
	Users []string
	Teams []string
}

type RequiredStatusCheck struct {
	Strict   bool
	Contexts []string
}

func FindProtections(ctx context.Context, owner string, repo string) ([]Protection, error) {
	protected := true
	opt := &github.BranchListOptions{
		Protected: &protected,
	}
	protectedBranches, _, err := ghc.Repositories.ListBranches(ctx, owner, repo, opt)
	if err != nil {
		return nil, err
	}

	protections := make([]Protection, len(protectedBranches))
	for i, pb := range protectedBranches {
		p, _, err := ghc.Repositories.GetBranchProtection(ctx, owner, repo, pb.GetName())
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
