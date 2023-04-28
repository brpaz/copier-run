package copier_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/brpaz/copier-run/internal/copier"
	"github.com/google/go-github/v52/github"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestServer(t *testing.T) *httptest.Server {
	router := mux.NewRouter()

	// Define the mock response for the /search/repositories endpoint
	router.HandleFunc("/search/repositories", func(w http.ResponseWriter, r *http.Request) {
		// Check the query parameter to make sure it matches our expectations
		q := r.URL.Query()
		assert.Equal(t, "topic:copier-template user:test-user", q.Get("q"))

		// Write the mock response
		resp := &github.RepositoriesSearchResult{
			Total: github.Int(1),
			Repositories: []*github.Repository{
				{
					Name:        github.String("test-repo"),
					Description: github.String("Test repo"),
					SSHURL:      github.String("git@github.com:test/test-repo.git"),
				},
			},
		}
		err := json.NewEncoder(w).Encode(resp)
		assert.NoError(t, err)
	})

	return httptest.NewServer(router)
}
func TestGitHubTemplateFetcher_Fetch(t *testing.T) {

	ts := createTestServer(t)
	defer ts.Close()

	// client is the GitHub client being tested and is
	// configured to use test server.
	client := github.NewClient(nil)

	uri, _ := url.Parse(ts.URL + "/")
	client.BaseURL = uri

	templateFetcher := copier.NewGitHubTemplateFetcher(client)

	// Call the Fetch method
	templates, err := templateFetcher.Fetch(context.Background(), "test-user")

	// Check that the Fetch method returns the expected result
	expectedTemplates := []copier.TemplateGitRepository{
		{
			Name:        "test-repo",
			Description: "Test repo",
			Url:         "git@github.com:test/test-repo.git",
		},
	}

	require.NoError(t, err)
	assert.Equal(t, expectedTemplates, templates)
}
