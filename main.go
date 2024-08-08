package main

import (
	"os"

	"github.com/juliflorezg/go-pocket-projects-gordle/gordle"
)

func main() {
	g := gordle.New(os.Stdin)

	g.Play()
}
