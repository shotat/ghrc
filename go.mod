module github.com/shotat/ghrc

go 1.13

replace github.com/google/go-github/v28 => github.com/shotat/go-github/v28 v28.0.3-alpha

require (
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/google/go-cmp v0.2.0
	github.com/google/go-github/v28 v28.0.1
	github.com/kr/pretty v0.1.0 // indirect
	github.com/spf13/cobra v0.0.5
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.2
)
