package spec

import "github.com/shotat/ghrc/state"

type Protection struct {
	Branch                     string                      `yaml:"branch"`
	RequiredStatusChecks       *RequiredStatusChecks       `yaml:"requiredStatusChecks,omitempty"`
	EnforceAdmins              *bool                       `yaml:"enforceAdmins,omitempty"`
	RequiredPullRequestReviews *RequiredPullRequestReviews `yaml:"requiredPullRequestReviews,omitempty"`
	Restrictions               *Restrictions               `yaml:"restrictions,omitempty"`
}

type Protections []Protection

type RequiredPullRequestReviews struct {
	DismissalRestrictions        Restrictions `yaml:"dismissalRestrictions"`
	DismissStaleReviews          bool         `yaml:"dismissStaleReviews"`
	RequireCodeOwnerReviews      bool         `yaml:"requireCodeOwnerReviews"`
	RequiredApprovingReviewCount int          `yaml:"requiredApprovingReviewCount"`
}

type Restrictions struct {
	Users []string `yaml:"users"`
	Teams []string `yaml:"teams"`
}

type RequiredStatusChecks struct {
	Strict   bool     `yaml:"strict"`
	Contexts []string `yaml:"contexts"`
}

func LoadProtectionsSpecFromState(states []state.Protection) Protections {
	specs := make([]Protection, len(states))
	for i, protection := range states {
		specs[i] = Protection{
			Branch:        protection.Branch,
			EnforceAdmins: protection.EnforceAdmins,
		}

		if protection.RequiredPullRequestReviews != nil {
			dismissalRestrictions := Restrictions{
				Users: protection.RequiredPullRequestReviews.DismissalRestrictions.Users,
				Teams: protection.RequiredPullRequestReviews.DismissalRestrictions.Teams,
			}

			specs[i].RequiredPullRequestReviews = &RequiredPullRequestReviews{
				DismissalRestrictions:        dismissalRestrictions,
				DismissStaleReviews:          protection.RequiredPullRequestReviews.DismissStaleReviews,
				RequireCodeOwnerReviews:      protection.RequiredPullRequestReviews.RequireCodeOwnerReviews,
				RequiredApprovingReviewCount: protection.RequiredPullRequestReviews.RequiredApprovingReviewCount,
			}
		}

		if protection.RequiredStatusChecks != nil {
			specs[i].RequiredStatusChecks = &RequiredStatusChecks{
				Strict:   protection.RequiredStatusChecks.Strict,
				Contexts: protection.RequiredStatusChecks.Contexts,
			}
		}

		if protection.Restrictions != nil {
			specs[i].Restrictions = &Restrictions{
				Users: protection.Restrictions.Users,
				Teams: protection.Restrictions.Teams,
			}
		}
	}
	return specs
}

// ToState returns a new state
func (sp *Protection) ToState() *state.Protection {
	newState := &state.Protection{}
	newState.Branch = sp.Branch

	if sp.RequiredStatusChecks != nil {
		newState.RequiredStatusChecks = &state.RequiredStatusChecks{}
		newState.RequiredStatusChecks.Strict = sp.RequiredStatusChecks.Strict
		newState.RequiredStatusChecks.Contexts = sp.RequiredStatusChecks.Contexts
	}

	if sp.EnforceAdmins != nil {
		newState.EnforceAdmins = sp.EnforceAdmins
	}

	if sp.RequiredPullRequestReviews != nil {
		newState.RequiredPullRequestReviews = &state.RequiredPullRequestReviews{}
		newState.RequiredPullRequestReviews.RequiredApprovingReviewCount = sp.RequiredPullRequestReviews.RequiredApprovingReviewCount
		newState.RequiredPullRequestReviews.DismissStaleReviews = sp.RequiredPullRequestReviews.DismissStaleReviews
		newState.RequiredPullRequestReviews.RequireCodeOwnerReviews = sp.RequiredPullRequestReviews.RequireCodeOwnerReviews
		newState.RequiredPullRequestReviews.DismissalRestrictions.Users = sp.RequiredPullRequestReviews.DismissalRestrictions.Users
		newState.RequiredPullRequestReviews.DismissalRestrictions.Teams = sp.RequiredPullRequestReviews.DismissalRestrictions.Teams
	}

	if sp.Restrictions != nil {
		newState.Restrictions = &state.Restrictions{}
		newState.Restrictions.Users = sp.Restrictions.Users
		newState.Restrictions.Teams = sp.Restrictions.Teams
	}
	return newState
}
