package gordle

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

var errInvalidWordLength = errors.New("invalid guess, word doesn't have the same number of characters as the solution")

// var errInvalidWordLength = fmt.Errorf("Invalid guess, word doesn't have the same number of characters as the solution")

func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	// return &Game{}

	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()
		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉You won! You found the word in %d attempt(s)! The words was: %s🎉\n", currentAttempt, string(g.solution))
			return
		}
	}
	fmt.Printf("😞 You've lost! The words was: %s\n", string(g.solution))

}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d character guess:\n", len(g.solution))

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle was not able to read your guess: %s\n", err.Error())
			continue
		}
		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.validateGuess(guess)

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s. \n", err.Error())
		} else {
			return guess
		}
	}
}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}

	return nil
}

func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}
