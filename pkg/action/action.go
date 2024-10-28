package action

import (
	"errors"
	"fmt"
	"path/filepath"
	"sync"

	gha "github.com/sethvargo/go-githubactions"

	"github.com/leonhfr/vocabulary-action/pkg/vocabulary"
)

const (
	languageInput   = "language"
	vocabularyInput = "vocabulary"
	directoryOutput = "directory"
	summaryOutput   = "summary"
)

func Run(action *gha.Action) error {
	input, err := newInput(action)
	if err != nil {
		return err
	}

	ctx := newContext(action)
	fh := &vocabulary.Filesystem{}
	parser := vocabulary.NewParser()

	config := vocabulary.ParseConfig(ctx.workspace)
	languageDir := vocabulary.LanguageDirectory(input.language, config)
	action.Infof(fmt.Sprintf("language \"%s\", target directory \"%s\"", input.language, languageDir))

	paragraphs := parser.Parse(input.vocabulary)
	accepted, discarded := vocabulary.Filter(paragraphs)
	for _, d := range discarded {
		action.Errorf("discarded \"%s\"", d)
	}
	buckets := vocabulary.Buckets(accepted)

	err = handle(action, ctx.workspace, languageDir, buckets, parser, fh)
	if err != nil {
		return err
	}

	setOutput(action, output{
		directory: languageDir,
		summary:   vocabulary.Summary(accepted),
	})

	return nil
}

func handle(action *gha.Action, workspace, languageDir string, buckets map[rune][]string, parser *vocabulary.Parser, fh vocabulary.FileHandler) error {
	targetDir := filepath.Join(workspace, languageDir)

	var errs []error
	wg := sync.WaitGroup{}

	for r, v := range buckets {
		wg.Add(1)

		go func(r rune, v []string) {
			defer wg.Done()

			filename := fmt.Sprintf("%c.md", r)
			action.Infof("adding vocabulary to file %s/%s", languageDir, filename)

			existing, err := vocabulary.Existing(targetDir, filename, fh)
			if err != nil {
				errs = append(errs, err)
				return
			}

			paragraphs := parser.Parse(existing)
			merged := vocabulary.Merge(paragraphs, v)
			err = vocabulary.Upsert(targetDir, filename, merged, fh)
			if err != nil {
				errs = append(errs, err)
			}
		}(r, v)
	}
	wg.Wait()

	if len(errs) > 0 {
		return fmt.Errorf("run: %w", errors.Join(errs...))
	}

	return nil
}

type input struct {
	language   string
	vocabulary string
}

func newInput(action *gha.Action) (input, error) {
	language := action.GetInput(languageInput)
	if language == "" {
		return input{}, errors.New("language cannot be an empty string")
	}

	vocabulary := action.GetInput(vocabularyInput)
	if vocabulary == "" {
		return input{}, errors.New("vocabulary cannot be an empty string")
	}

	return input{
		language:   language,
		vocabulary: vocabulary,
	}, nil
}

type context struct {
	workspace string
}

func newContext(action *gha.Action) context {
	c, _ := action.Context()
	return context{workspace: c.Workspace}
}

type output struct {
	directory string
	summary   string
}

func setOutput(action *gha.Action, output output) {
	action.SetOutput(directoryOutput, output.directory)
	action.SetOutput(summaryOutput, output.summary)
}
