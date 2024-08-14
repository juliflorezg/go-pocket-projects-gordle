package gordle

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader *bufio.Reader
}

const solutionLength = 5

var errInvalidWordLength = errors.New("invalid guess, word doesn't have the same number of characters as the solution")

// var errInvalidWordLength = fmt.Errorf("Invalid guess, word doesn't have the same number of characters as the solution")

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

	// fmt.Printf("Enter a guess: \n")

	guess := g.ask()

	fmt.Printf("Your guess is: %s\n", string(guess))

}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle was not able to read your guess: %s\n", err.Error())
			continue
		}
		guess := []rune(string(playerInput))

		err = g.validateGuess(guess)

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s. \n", err.Error())
		} else {
			return guess
		}
	}
}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d, got %d, %w", solutionLength, len(guess), errInvalidWordLength)
	}

	return nil
}
