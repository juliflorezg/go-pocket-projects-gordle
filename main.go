package main

import (
	"os"

	"github.com/juliflorezg/go-pocket-projects-gordle/gordle"
)

const maxAttempts = 6

func main() {
	g := gordle.New(os.Stdin, "skates", maxAttempts)

	g.Play()
}
