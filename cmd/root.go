package cmd

import (
	"context"

	"github.com/brpaz/copier-run/internal/copier"
	"github.com/brpaz/copier-run/internal/github"
	"github.com/brpaz/copier-run/internal/prompt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var githubUser = "brpaz"

// NewRootCmd returns a new instance of the root command for the application
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "copier-run",
		Short: "Provides a list of available copier templates for a GitHub user and allows to execute them",
		RunE: func(cmd *cobra.Command, args []string) error {

			ctx := context.Background()

			gitHubClient := github.NewClient(ctx)
			templateFetcher := copier.NewGitHubTemplateFetcher(gitHubClient)

			templates, err := templateFetcher.Fetch(ctx, githubUser)

			if err != nil {
				return errors.Wrap(err, "Error fetching templates from GitHub")
			}

			answers, err := prompt.Execute(templates)

			if err != nil {
				return errors.Wrap(err, "Error selecting a template")
			}

			if answers.SelectedTemplate != nil {
				err := copier.RunGenerator(answers.SelectedTemplate.Url, answers.SelectedDestination)

				if err != nil {
					return errors.Wrap(err, "Error running copier-runerator")
				}
			}

			return nil
		},
	}

	rootCmd.Flags().StringVar(&githubUser, "gh-user", githubUser, "GitHub user to fetch the templates from")

	// Reggister subcommands
	rootCmd.AddCommand(NewVersionCmd())

	return rootCmd
}
