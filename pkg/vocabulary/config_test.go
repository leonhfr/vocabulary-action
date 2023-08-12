package vocabulary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_languageDirectory(t *testing.T) {
	tests := []struct {
		name     string
		cfg      config
		language string
		want     string
	}{
		{
			"in config",
			config{
				"de": "german/vocabulary",
			},
			"de",
			"german/vocabulary",
		},
		{
			"in config",
			config{
				"default": "random/vocabulary",
			},
			"de",
			"random/vocabulary",
		},
		{
			"in config",
			config{},
			"de",
			"todo/vocabulary",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cfg.languageDirectory(tt.language)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_newConfig(t *testing.T) {
	cfg, err := newConfig("../../test/data")
	assert.Equal(t, config{
		"ca":      "catalan/vocabulary",
		"de":      "german/vocabulary",
		"en":      "english/vocabulary",
		"es":      "spanish/vocabulary",
		"default": "todo/vocabulary",
	}, cfg)
	assert.NoError(t, err)
}
