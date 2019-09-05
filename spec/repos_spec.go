package spec

import (
	"github.com/shotat/ghrc/status"
)

type RepositorySpec struct {
	Description      *string `yaml:"description,omitempty"`
	Homepage         *string `yaml:"homepage,omitempty"`
	Private          *bool   `yaml:"private"`
	AllowSquashMerge *bool   `yaml:"allowSquashMerge"`
	AllowMergeCommit *bool   `yaml:"allowMergeCommit"`
	AllowRebaseMerge *bool   `yaml:"allowRebaseMerge"`

	Topics      []string     `yaml:"topics,omitempty"`
	Labels      []Label      `yaml:"labels,omitempty"`
	Protections []Protection `yaml:"protections,omitempty"`
}

func (sp *RepositorySpec) Patch(st *status.RepositoryStatus) {
	if sp.Description != nil {
		st.Description = sp.Description
	}
	if sp.Private != nil {
		st.Private = sp.Private
	}
	if sp.Homepage != nil {
		st.Homepage = sp.Homepage
	}
	if sp.AllowSquashMerge != nil {
		st.AllowSquashMerge = sp.AllowSquashMerge
	}
	if sp.AllowMergeCommit != nil {
		st.AllowMergeCommit = sp.AllowMergeCommit
	}
	if sp.AllowRebaseMerge != nil {
		st.AllowRebaseMerge = sp.AllowRebaseMerge
	}
	if sp.Topics != nil {
		st.Topics = sp.Topics
	}
	if sp.Labels != nil {
		labels := make([]status.Label, len(sp.Labels))
		for i, spl := range sp.Labels {
			labels[i] = func() status.Label {
				for _, stl := range st.Labels {
					if stl.Name == spl.Name {
						// update existing label
						stl.Name = spl.Name
						stl.Color = spl.Color
						if spl.Description != nil {
							stl.Description = spl.Description
						}
						return stl
					}
				}
				// new label
				return status.Label{
					ID:          nil,
					Name:        spl.Name,
					Color:       spl.Color,
					Description: spl.Description,
				}
			}()
		}
		st.Labels = labels
	}

	if sp.Protections != nil {
		protections := make([]status.Protection, len(sp.Protections))
		for i, spp := range sp.Protections {
			protections[i] = func() status.Protection {
				for _, stp := range st.Protections {
					if stp.Branch == spp.Branch {
						if spp.EnforceAdmins != nil {
							stp.EnforceAdmins = spp.EnforceAdmins
							return stp
						}
					}
				}

				// new protection
				return status.Protection{
					Branch:        spp.Branch,
					EnforceAdmins: spp.EnforceAdmins,
					// RequiredStatusCheck:        spp.RequiredStatusCheck,
					// RequiredPullRequestReviews: *spp.RequiredPullRequestReviews,
				}
			}()
		}
		st.Protections = protections
	}
}

type Label struct {
	Name        string  `yaml:"name"`
	Description *string `yaml:"description,omitempty"`
	Color       string  `yaml:"color"`
}

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
