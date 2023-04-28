package copier

import (
	"context"
	"fmt"

	"github.com/google/go-github/v52/github"
	"github.com/pkg/errors"
)

// TemplateGitRepository represents a copier template repository
type TemplateGitRepository struct {
	Name        string
	Description string
	Url         string
}

type TemplateFetcher interface {
	Fetch() ([]TemplateGitRepository, error)
}

type GitHubTemplateFetcher struct {
	gitHub *github.Client
}

func NewGitHubTemplateFetcher(client *github.Client) *GitHubTemplateFetcher {
	return &GitHubTemplateFetcher{gitHub: client}
}

// Fetch returns a list of available copier templates for the given user.
// It identifies the repositories by checking the "topic" attribute matches "copier-template"
func (tf GitHubTemplateFetcher) Fetch(ctx context.Context, user string) ([]TemplateGitRepository, error) {

	topic := "copier-template"
	templates := make([]TemplateGitRepository, 0)

	result, _, err := tf.gitHub.Search.Repositories(ctx, fmt.Sprintf("topic:%s user:%s", topic, user), &github.SearchOptions{
		Sort: "name",
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	})

	if err != nil {
		return templates, errors.Wrap(err, "Error fetching repositories from GitHub")
	}

	for _, repo := range result.Repositories {
		templates = append(templates, TemplateGitRepository{
			Name:        repo.GetName(),
			Description: repo.GetDescription(),
			Url:         repo.GetSSHURL(),
		})
	}

	return templates, nil
}
