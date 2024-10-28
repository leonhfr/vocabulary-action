package vocabulary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LanguageDirectory(t *testing.T) {
	tests := []struct {
		name     string
		config   map[string]string
		language string
		want     string
	}{
		{
			"in config",
			map[string]string{
				"de": "german/vocabulary",
			},
			"de",
			"german/vocabulary",
		},
		{
			"in config",
			map[string]string{
				"default": "random/vocabulary",
			},
			"de",
			"random/vocabulary",
		},
		{
			"in config",
			map[string]string{},
			"de",
			"todo/vocabulary",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LanguageDirectory(tt.language, tt.config)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_ParseConfig(t *testing.T) {
	config := ParseConfig("../../test/data")
	assert.Equal(t, map[string]string{
		"ca":      "catalan/vocabulary",
		"de":      "german/vocabulary",
		"en":      "english/vocabulary",
		"es":      "spanish/vocabulary",
		"default": "todo/vocabulary",
	}, config)
}

func Test_Buckets(t *testing.T) {
	vocabulary := []string{
		"Spaziergang",
		"Spiegel",
		"Fernweh",
	}
	got := Buckets(vocabulary)
	assert.Equal(t, map[rune][]string{
		's': {"Spaziergang", "Spiegel"},
		'f': {"Fernweh"},
	}, got)
}

func Test_Summary(t *testing.T) {
	vocabulary := []string{
		"süchtig \"süchtig machen \"",
		"Spaziergang",
		"Spiegel",
		"Schuh #cognate",
	}
	got := Summary(vocabulary)
	assert.Equal(t, "Schuh, Spaziergang, Spiegel, süchtig", got)
}
