package main

import (
	"github.com/sethvargo/go-githubactions"

	"github.com/leonhfr/vocabulary-action/pkg/action"
	"github.com/leonhfr/vocabulary-action/pkg/vocabulary"
)

func main() {
	a := githubactions.New()
	h := vocabulary.Handler{}

	if err := action.Run(a, h); err != nil {
		a.Fatalf("%v", err)
	}
}
