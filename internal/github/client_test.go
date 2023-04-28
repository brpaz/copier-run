package github

import (
	"context"
	"testing"

	"github.com/google/go-github/v52/github"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestNewClient(t *testing.T) {
	// Set up test context
	ctx := context.Background()

	// Set up mock token
	token := "mock-token"

	t.Setenv("GITHUB_TOKEN", token)

	// Set up expected client
	expectedClient := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)))

	// Call NewClient function
	client := NewClient(ctx)

	// Assert that the returned client matches the expected client
	assert.Equal(t, expectedClient, client)
}
