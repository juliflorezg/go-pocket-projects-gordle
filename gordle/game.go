package gordle

import (
	"bufio"
	"fmt"
	"io"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader *bufio.Reader
}

func New(playerInput io.Reader) *Game {
	// return &Game{}

	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	fmt.Printf("Enter a guess: \n")

}
