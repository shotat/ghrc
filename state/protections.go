package state

import (
	"context"
	"github.com/google/go-github/v28/github"
)

type Protection struct {
	Branch                     string
	RequiredStatusChecks       RequiredStatusChecks
	EnforceAdmins              bool
	RequiredPullRequestReviews RequiredPullRequestReviews
	Restrictions               Restrictions
}

type RequiredPullRequestReviews struct {
	DismissalRestrictions        Restrictions
	DismissStaleReviews          bool
	RequireCodeOwnerReviews      bool
	RequiredApprovingReviewCount int
}

type Restrictions struct {
	Users []string
	Teams []string
}

type RequiredStatusChecks struct {
	Strict   bool
	Contexts []string
}

func (p *Protection) Update(ctx context.Context, repoOwner string, repoName string) error {
	req := &github.ProtectionRequest{
		// RequiredStatusChecks:       p.RequiredStatusChecks,
		// RequiredPullRequestReviews: p.RequiredPullRequestReviews,
		EnforceAdmins: p.EnforceAdmins,
		// Restrictions:               p.Restrictions,
	}
	_, _, err := ghc.Repositories.UpdateBranchProtection(ctx, repoOwner, repoName, p.Branch, req)
	return err
}

func (p *Protection) Destroy(ctx context.Context, repoOwner string, repoName string) error {
	_, err := ghc.Repositories.RemoveBranchProtection(ctx, repoOwner, repoName, p.Branch)
	return err
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

		dismissalRestrictions := Restrictions{
			Users: []string{},
			Teams: []string{},
		}
		for _, u := range p.GetRequiredPullRequestReviews().DismissalRestrictions.Users {
			dismissalRestrictions.Users = append(dismissalRestrictions.Users, u.GetLogin())
		}
		for _, t := range p.GetRequiredPullRequestReviews().DismissalRestrictions.Teams {
			dismissalRestrictions.Teams = append(dismissalRestrictions.Teams, t.GetSlug())
		}

		restrictions := Restrictions{
			Users: []string{},
			Teams: []string{},
		}
		for _, u := range p.GetRestrictions().Users {
			restrictions.Users = append(restrictions.Users, u.GetLogin())
		}
		for _, t := range p.GetRestrictions().Teams {
			restrictions.Teams = append(restrictions.Teams, t.GetSlug())
		}
		protections[i] = Protection{
			Branch: pb.GetName(),
			RequiredStatusChecks: RequiredStatusChecks{
				Strict:   p.GetRequiredStatusChecks().Strict,
				Contexts: p.GetRequiredStatusChecks().Contexts,
			},
			EnforceAdmins: p.GetEnforceAdmins().Enabled,
			RequiredPullRequestReviews: RequiredPullRequestReviews{
				DismissalRestrictions:        dismissalRestrictions,
				DismissStaleReviews:          p.GetRequiredPullRequestReviews().DismissStaleReviews,
				RequireCodeOwnerReviews:      p.GetRequiredPullRequestReviews().RequireCodeOwnerReviews,
				RequiredApprovingReviewCount: p.GetRequiredPullRequestReviews().RequiredApprovingReviewCount,
			},
			Restrictions: restrictions,
		}
	}
	return protections, err
}
