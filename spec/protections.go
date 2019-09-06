package spec

import (
	"github.com/shotat/ghrc/change"
	"github.com/shotat/ghrc/state"
)

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

func (sp *Repo) GetProtectionsChange(st []state.Protection) *change.ReposChange {
	/*
		if sp.Protections != nil {
			protectionsPatch := new(state.BulkPatch)
			protections := make([]state.Protection, len(sp.Protections))
			for i, spp := range sp.Protections {
				protections[i] = func() state.Protection {
					for _, stp := range st.Protections {
						if stp.Branch == spp.Branch {
							if spp.EnforceAdmins != nil {
								stp.EnforceAdmins = spp.EnforceAdmins
								return stp
							}
						}
					}

					// new protection
					return state.Protection{
						Branch:        spp.Branch,
						EnforceAdmins: spp.EnforceAdmins,
						// RequiredStatusCheck:        spp.RequiredStatusCheck,
						// RequiredPullRequestReviews: *spp.RequiredPullRequestReviews,
					}
				}()
			}
			st.Protections = protections
		}
	*/
	return nil
}
