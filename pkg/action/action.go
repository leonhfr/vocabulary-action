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

	output, err := runner.Run(input, newLogger(action), &vocabulary.Filesystem{})
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

type logger struct {
	action *githubactions.Action
}

func newLogger(action *githubactions.Action) *logger {
	return &logger{action}
}

func (l *logger) Info(msg string, args ...any) {
	l.action.Infof(msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.action.Warningf(msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.action.Errorf(msg, args...)
}
