module github.com/shotat/ghrc

go 1.13

replace github.com/google/go-github/v28 => github.com/shotat/go-github/v28 v28.0.3-alpha

require (
	github.com/google/go-github/v28 v28.0.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/yaml.v2 v2.2.2
)
