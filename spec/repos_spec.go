package spec

import (
	"github.com/shotat/ghrc/change"
	"github.com/shotat/ghrc/status"
)

type Repo struct {
	Description      *string `yaml:"description,omitempty"`
	Homepage         *string `yaml:"homepage,omitempty"`
	Private          *bool   `yaml:"private"`
	AllowSquashMerge *bool   `yaml:"allowSquashMerge"`
	AllowMergeCommit *bool   `yaml:"allowMergeCommit"`
	AllowRebaseMerge *bool   `yaml:"allowRebaseMerge"`

	Topics []string `yaml:"topics,omitempty"`
}

type Spec struct {
	Repo        *Repo        `yaml:"repo,omitempty"`
	Labels      []Label      `yaml:"labels,omitempty"`
	Protections []Protection `yaml:"protections,omitempty"`
}

func GetLabelsChangeSet(st []status.Label, sp []Label) []change.LabelChange {
	if sp == nil {
		return nil
	}
	changes := make([]change.LabelChange, 0)
	for _, spl := range sp {
		func() {
			for _, stl := range st {
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
					changes = append(changes, change.LabelChange{
						Action: change.Update,
						Before: &stl,
						After:  &after,
					})
					return
				}
			}
			// new label
			changes = append(changes, change.LabelChange{
				Action: change.Create,
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
	for _, stl := range st {
		func() {
			for _, spl := range sp {
				if stl.Name == spl.Name {
					return
				}
			}

			// deletion
			changes = append(changes, change.LabelChange{
				Action: change.Delete,
				Before: &stl,
				After:  nil,
			})
			return
		}()
	}
	return changes
}
func (sp *Repo) GetRepoChange(st *status.Repo) *change.ReposChange {
	after := &status.Repo{
		Description:      st.Description,
		Homepage:         st.Homepage,
		Private:          st.Private,
		AllowSquashMerge: st.AllowSquashMerge,
		AllowMergeCommit: st.AllowMergeCommit,
		AllowRebaseMerge: st.AllowRebaseMerge,
		Topics:           st.Topics,
	}

	if sp.Description != nil {
		after.Description = sp.Description
	}
	if sp.Private != nil {
		after.Private = sp.Private
	}
	if sp.Homepage != nil {
		after.Homepage = sp.Homepage
	}
	if sp.AllowSquashMerge != nil {
		after.AllowSquashMerge = sp.AllowSquashMerge
	}
	if sp.AllowMergeCommit != nil {
		after.AllowMergeCommit = sp.AllowMergeCommit
	}
	if sp.AllowRebaseMerge != nil {
		after.AllowRebaseMerge = sp.AllowRebaseMerge
	}
	if sp.Topics != nil {
		after.Topics = sp.Topics
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
	return &change.ReposChange{
		Action: change.Update,
		Before: st,
		After:  after,
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
