package gordle

import "fmt"

// Game holds all the information we need to play a game of gordle.
type Game struct {
}

func New() *Game {
	// return &Game{}

	g := &Game{}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	fmt.Printf("Enter a guess: \n")

}
