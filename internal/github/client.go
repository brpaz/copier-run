package github

import (
	"context"
	"os"

	goGitHub "github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
)

// NewClient creates a new instance of the GitHub client
func NewClient(ctx context.Context) *goGitHub.Client {
	token := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	return goGitHub.NewClient(tc)
}
