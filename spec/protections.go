package spec

import "github.com/shotat/ghrc/state"

type Protection struct {
	Branch                     *string                     `yaml:"branch"`
	RequiredStatusCheck        *RequiredStatusCheck        `yaml:"requiredStatusCheck"`
	EnforceAdmins              *bool                       `yaml:"enforceAdmins"`
	RequiredPullRequestReviews *RequiredPullRequestReviews `yaml:"requiredPullRequestReviews"`
	Restrictions               *Restrictions               `yaml:"restrictions"`
}

type Protections []Protection

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

// ToState merge state and spec to generate new state
func (sp *Protection) ToState(base *state.Protection) *state.Protection {
	newState := &state.Protection{}
	// initialize
	if base != nil {
		newState.Branch = base.Branch
		newState.RequiredStatusCheck = base.RequiredStatusCheck
		newState.EnforceAdmins = base.EnforceAdmins
		newState.RequiredPullRequestReviews = base.RequiredPullRequestReviews
		newState.Restrictions = base.Restrictions
	}
	newState.EnforceAdmins = sp.EnforceAdmins
	// TODO
	return newState
}
