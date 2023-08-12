package action

import (
	"errors"

	"github.com/sethvargo/go-githubactions"

	"github.com/leonhfr/vocabulary-action/pkg/vocabulary"
)

const (
	languageInput   = "language"
	vocabularyInput = "vocabulary"
	directoryOutput = "directory"
	summaryOutput   = "summary"
)

func Run(action *githubactions.Action, runner vocabulary.Runner) error {
	input, err := newInput(action)
	if err != nil {
		return err
	}

	output, err := runner.Run(input)
	if err != nil {
		return err
	}

	action.SetOutput(directoryOutput, output.Directory)
	action.SetOutput(summaryOutput, output.Summary)

	return nil
}

func newInput(action *githubactions.Action) (vocabulary.Input, error) {
	l := action.GetInput(languageInput)
	if l == "" {
		return vocabulary.Input{}, errors.New("language cannot be an empty string")
	}

	v := action.GetInput(vocabularyInput)
	if v == "" {
		return vocabulary.Input{}, errors.New("vocabulary cannot be an empty string")
	}

	c, err := action.Context()
	if err != nil {
		return vocabulary.Input{}, err
	}

	return vocabulary.Input{
		Language:   l,
		Vocabulary: v,
		Workspace:  c.Workspace,
	}, nil
}
