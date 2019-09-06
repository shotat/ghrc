package metadata

type Metadata struct {
	Owner string `yaml:"owner"`
	Name  string `yaml:"name"`
}

func NewMetadata(repoOwner string, repoName string) *Metadata {
	return &Metadata{
		Owner: repoOwner,
		Name:  repoName,
	}
}
