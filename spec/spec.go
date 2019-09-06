package spec

type Spec struct {
	Repo        *Repo       `yaml:"repo,omitempty"`
	Labels      Labels      `yaml:"labels,omitempty"`
	Protections Protections `yaml:"protections,omitempty"`
}
