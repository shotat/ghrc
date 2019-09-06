package change

import (
	"github.com/shotat/ghrc/spec"
	"github.com/shotat/ghrc/state"
)

func GetProtectionsChange(st []state.Protection, sp *spec.Repo) *ReposChange {
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
