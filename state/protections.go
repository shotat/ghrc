package state

import (
	"context"
	"github.com/google/go-github/v28/github"
)

type Protection struct {
	Branch                     string
	RequiredStatusChecks       *RequiredStatusChecks
	EnforceAdmins              *bool
	RequiredPullRequestReviews *RequiredPullRequestReviews
	Restrictions               *Restrictions
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
	requiredStatusChecks := &github.RequiredStatusChecks{
		Strict:   p.RequiredStatusChecks.Strict,
		Contexts: p.RequiredStatusChecks.Contexts,
	}

	restrictions := &github.BranchRestrictionsRequest{
		Users: []string{},
		Teams: []string{},
	}
	for _, u := range p.Restrictions.Users {
		restrictions.Users = append(restrictions.Users, u)
	}
	for _, t := range p.Restrictions.Teams {
		restrictions.Teams = append(restrictions.Teams, t)
	}
	dismissalRestrictions := &github.DismissalRestrictionsRequest{
		Users: &[]string{},
		Teams: &[]string{},
	}
	for _, u := range p.RequiredPullRequestReviews.DismissalRestrictions.Users {
		tmp := append(*dismissalRestrictions.Users, u)
		dismissalRestrictions.Users = &tmp
	}
	for _, t := range p.RequiredPullRequestReviews.DismissalRestrictions.Teams {
		tmp := append(*dismissalRestrictions.Teams, t)
		dismissalRestrictions.Teams = &tmp
	}

	// FIXME: only org
	_ = dismissalRestrictions
	_ = restrictions
	requiredPullRequestReviews := &github.PullRequestReviewsEnforcementRequest{
		// DismissalRestrictionsRequest: dismissalRestrictions,
		DismissStaleReviews:          p.RequiredPullRequestReviews.DismissStaleReviews,
		RequireCodeOwnerReviews:      p.RequiredPullRequestReviews.RequireCodeOwnerReviews,
		RequiredApprovingReviewCount: p.RequiredPullRequestReviews.RequiredApprovingReviewCount,
	}

	req := &github.ProtectionRequest{
		RequiredStatusChecks:       requiredStatusChecks,
		RequiredPullRequestReviews: requiredPullRequestReviews,
		// Restrictions:               restrictions,
	}
	if p.EnforceAdmins != nil {
		req.EnforceAdmins = *p.EnforceAdmins
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

		protections[i] = Protection{
			Branch:        pb.GetName(),
			EnforceAdmins: &p.GetEnforceAdmins().Enabled,
		}

		if p.GetRestrictions() != nil {
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
			protections[i].Restrictions = &restrictions
		}

		if p.GetRequiredStatusChecks() != nil {
			protections[i].RequiredStatusChecks = &RequiredStatusChecks{
				Strict:   p.GetRequiredStatusChecks().Strict,
				Contexts: p.GetRequiredStatusChecks().Contexts,
			}
		}

		if p.GetRequiredPullRequestReviews() != nil {
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
			protections[i].RequiredPullRequestReviews = &RequiredPullRequestReviews{
				DismissalRestrictions:        dismissalRestrictions,
				DismissStaleReviews:          p.GetRequiredPullRequestReviews().DismissStaleReviews,
				RequireCodeOwnerReviews:      p.GetRequiredPullRequestReviews().RequireCodeOwnerReviews,
				RequiredApprovingReviewCount: p.GetRequiredPullRequestReviews().RequiredApprovingReviewCount,
			}
		}
	}
	return protections, err
}
