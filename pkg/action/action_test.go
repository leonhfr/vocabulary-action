package action

import (
	"io"
	"testing"

	"github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/assert"

	"github.com/leonhfr/vocabulary-action/pkg/vocabulary"
)

var _ vocabulary.Logger = &logger{}

func Test_newInput(t *testing.T) {
	l, v, w := "de", "Spaziergang", "/vocabulary"

	tests := []struct {
		name   string
		envMap map[string]string
		want   vocabulary.Input
		err    string
	}{
		{
			name: "working",
			envMap: map[string]string{
				"INPUT_LANGUAGE":   l,
				"INPUT_VOCABULARY": v,
				"GITHUB_WORKSPACE": w,
			},
			want: vocabulary.Input{
				Language:   l,
				Vocabulary: v,
				Workspace:  w,
			},
			err: "",
		},
		{
			name: "missing language",
			envMap: map[string]string{
				"INPUT_VOCABULARY": v,
				"GITHUB_WORKSPACE": w,
			},
			want: vocabulary.Input{},
			err:  "language cannot be an empty string",
		},
		{
			name: "missing language",
			envMap: map[string]string{
				"INPUT_LANGUAGE":   l,
				"GITHUB_WORKSPACE": w,
			},
			want: vocabulary.Input{},
			err:  "vocabulary cannot be an empty string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getenv := func(key string) string {
				return tt.envMap[key]
			}

			action := githubactions.New(
				githubactions.WithWriter(io.Discard),
				githubactions.WithGetenv(getenv),
			)

			got, err := newInput(action)

			assert.Equal(t, tt.want, got)
			if tt.err == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}
