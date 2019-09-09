package spec

import (
	"github.com/shotat/ghrc/state"
)

type Repo struct {
	Description      string `yaml:"description"`
	Homepage         string `yaml:"homepage"`
	Private          bool   `yaml:"private"`
	AllowSquashMerge bool   `yaml:"allowSquashMerge"`
	AllowMergeCommit bool   `yaml:"allowMergeCommit"`
	AllowRebaseMerge bool   `yaml:"allowRebaseMerge"`

	Topics []string `yaml:"topics"`
}

func LoadRepoSpecFromState(st *state.Repo) *Repo {
	return &Repo{
		Homepage:         st.Homepage,
		Description:      st.Description,
		Private:          st.Private,
		Topics:           st.Topics,
		AllowSquashMerge: st.AllowSquashMerge,
		AllowMergeCommit: st.AllowMergeCommit,
		AllowRebaseMerge: st.AllowRebaseMerge,
	}
}
