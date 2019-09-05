package spec

import (
	"github.com/shotat/ghrc/patch"
	"github.com/shotat/ghrc/status"
)

/*

type EntirePatch struct  {
	RepositoryPatch RepositoryPatch
	labelPatches []LabelPatch
	branchPatches []BranchPatch

	patch.String()
	CHANGE "foo" => "bar"
	DELETE "ready_for_review" =
	CHANGE "foo" => "bar"
	type Delelte struct {
		Resource
	}
	type Create struct {
		Resource
	}
	type Change struct {
		Old Resource
		New Resource
	}
}

*/

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

type RepositoryPatch struct {
	Description      *string
	Homepage         *string
	Private          *bool
	AllowSquashMerge *bool
	AllowMergeCommit *bool
	AllowRebaseMerge *bool
	// ProtectionsPatch *status.BulkPatch
	LabelPatches []patch.LabelPatch
}

func (sp *RepositorySpec) CalculatePatch(st *status.RepositoryStatus) *RepositoryPatch {
	p := new(RepositoryPatch)
	if sp.Description != nil && st.Description != sp.Description {
		p.Description = sp.Description
	}
	if sp.Private != nil && st.Private != sp.Private {
		p.Private = sp.Private
	}
	if sp.Homepage != nil && st.Homepage != sp.Homepage {
		p.Homepage = sp.Homepage
	}
	if sp.AllowSquashMerge != nil && st.AllowSquashMerge != sp.AllowSquashMerge {
		p.AllowSquashMerge = sp.AllowSquashMerge
	}
	/*
		if sp.AllowMergeCommit != nil {
			st.AllowMergeCommit != sp.AllowMergeCommit
		}
		if sp.AllowRebaseMerge != nil {
			st.AllowRebaseMerge != sp.AllowRebaseMerge
		}
		if sp.Topics != nil {
			st.Topics = sp.Topics
		}
	*/
	if sp.Labels != nil {
		labelPatches := make([]patch.LabelPatch, 0)
		for _, spl := range sp.Labels {
			func() {
				for _, stl := range st.Labels {
					if stl.Name == spl.Name {
						// update existing label
						after := status.Label{
							Name:        spl.Name,
							Color:       spl.Color,
							Description: stl.Description,
						}
						if spl.Description != nil {
							after.Description = spl.Description
						}
						labelPatches = append(labelPatches, patch.LabelPatch{
							Before: &stl,
							After:  &after,
						})
						return
					}
				}
				// new label
				labelPatches = append(labelPatches, patch.LabelPatch{
					Before: nil,
					After: &status.Label{
						Name:        spl.Name,
						Color:       spl.Color,
						Description: spl.Description,
					},
				})
				return
			}()
		}
		for _, stl := range st.Labels {
			func() {
				for _, spl := range sp.Labels {
					if stl.Name == spl.Name {
						return
					}
				}

				// deletion
				labelPatches = append(labelPatches, patch.LabelPatch{
					Before: &stl,
					After:  nil,
				})
				return
			}()
		}
		p.LabelPatches = labelPatches
	}

	/*
		if sp.Protections != nil {
			protectionsPatch := new(status.BulkPatch)
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
	*/
	return p
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
