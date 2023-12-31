package vocabulary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	test := `<!-- This file was generated by github.com/leonhfr/vocabulary-action and is susceptible to be modified by automations. -->

# s

Schuh

Spaziergang "Example phrase." "Another example."

Spiegel
A note about the word.
Can be multiline.
`

	expected := []string{
		"Schuh",
		"Spaziergang \"Example phrase.\" \"Another example.\"",
		"Spiegel\nA note about the word.\nCan be multiline.",
	}

	p := NewParser()
	vocabulary := p.Parse(test)

	assert.Equal(t, expected, vocabulary)
}
