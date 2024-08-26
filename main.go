package main

import (
	"fmt"
	"os"

	"github.com/juliflorezg/go-pocket-projects-gordle/gordle"
)

const maxAttempts = 6

func main() {
	corpus, err := gordle.ReadCorpus("corpus/english.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "encountered an error opening a words file: %v\n", err)
		return
	}

	g, err := gordle.New(os.Stdin, corpus, maxAttempts)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start game: %s", err)
		return
	}

	g.Play()
}
