module github.com/shotat/ghrc

go 1.12

replace github.com/google/go-github/v28 => github.com/shotat/go-github/v28 v28.0.2-0.20190904052026-45425c447656

require (
	github.com/google/go-github/v28 v28.0.1
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
	gopkg.in/yaml.v2 v2.2.2
)
