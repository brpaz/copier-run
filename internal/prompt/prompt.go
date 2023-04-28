package prompt

import (
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/brpaz/copier-run/internal/copier"
	"github.com/pkg/errors"
)

type PromptResult struct {
	SelectedTemplate    *copier.TemplateGitRepository
	SelectedDestination string
}

func getAutocompleteSuggestions(baseDirectory string) []string {

	matches, _ := filepath.Glob(baseDirectory + "*")

	var dirs []string

	for _, match := range matches {
		fileInfo, err := os.Stat(match)

		if err != nil {
			continue
		}

		if fileInfo.IsDir() {
			dirs = append(dirs, match)
		}
	}
	return dirs
}

func buildQuestionsPrompt(availableTemplates []copier.TemplateGitRepository) []*survey.Question {

	availableTemplateNames := make([]string, len(availableTemplates))

	for i, template := range availableTemplates {
		availableTemplateNames[i] = template.Name
	}

	questions := []*survey.Question{
		{
			Name: "template",
			Prompt: &survey.Select{
				Message: "Choose a template:",
				Options: availableTemplateNames,
			},
			Validate: survey.Required,
		},
		{
			Name: "destination",
			Prompt: &survey.Input{
				Message: "Choose a destination directory:",
				Suggest: func(toComplete string) []string {
					return getAutocompleteSuggestions(toComplete)
				},
			},
			Validate: survey.Required,
		},
	}

	return questions

}

// Execute executes the prompt that asks the user for the template and destination directory
func Execute(availableTemplates []copier.TemplateGitRepository) (PromptResult, error) {

	qs := buildQuestionsPrompt(availableTemplates)

	// the answers will be written to this struct
	answers := struct {
		Template    string `survey:"template"`
		Destination string `survey:"destination"`
	}{}

	err := survey.Ask(
		qs,
		&answers,
	)

	if err != nil {
		return PromptResult{}, errors.Wrap(err, "Error processing answers")
	}

	var selectedTemplateItem copier.TemplateGitRepository

	for _, template := range availableTemplates {
		if template.Name == answers.Template {
			selectedTemplateItem = template
			break
		}
	}

	return PromptResult{
		SelectedTemplate:    &selectedTemplateItem,
		SelectedDestination: answers.Destination,
	}, nil
}
