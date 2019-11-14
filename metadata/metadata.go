package metadata

type Metadata struct {
	Name string `yaml:"name"`
}

func NewMetadata(repoName string) *Metadata {
	return &Metadata{
		Name: repoName,
	}
}
