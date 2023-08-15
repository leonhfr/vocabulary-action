package main

import (
	"github.com/sethvargo/go-githubactions"

	"github.com/leonhfr/vocabulary-action/pkg/action"
)

func main() {
	gha := githubactions.New()

	if err := action.Run(gha); err != nil {
		gha.Fatalf("%v", err)
	}
}
