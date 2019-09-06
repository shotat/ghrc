package state

import (
	"context"
	"os"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
	"net/url"
)

var ghc *github.Client

func init() {
	accessToken := os.Getenv("GHRC_GITHUB_TOKEN")
	api := os.Getenv("GHRC_GITHUB_API")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	// initialize ghc var
	ghc = github.NewClient(tc)

	if api != "" {
		parsedURL, err := url.Parse(api)
		if err != nil {
			panic(err)
		}
		ghc.BaseURL = parsedURL
		ghc.UploadURL = parsedURL
	}
}
