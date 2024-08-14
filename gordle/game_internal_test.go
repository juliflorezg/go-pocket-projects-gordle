package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {

	testCases := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "ﺎﺒﺣﺮﻣ",
			want:  []rune("ﺎﺒﺣﺮﻣ"),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input), tc.input, 0)

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v | want =%v ", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		input []rune
		want  error
	}{
		"correct length": {
			input: []rune("hello"),
			want:  nil,
		},
		"length too short": {
			input: []rune("hel"),
			want:  errInvalidWordLength,
		},
		"length too long": {
			input: []rune("hellooo"),
			want:  errInvalidWordLength,
		},
		"empty input": {
			input: []rune(""),
			want:  errInvalidWordLength,
		},
		"no input": {
			input: nil,
			want:  errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(string(tc.input)), string(tc.input), 0)

			got := g.validateGuess(tc.input)

			if !errors.Is(got, tc.want) {
				t.Errorf("word: %c => got: %v | want: %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestSplitToUppercaseCharacters(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"all uppercase": {
			input: "UPPER",
			want:  []rune("UPPER"),
		},
		"all lowercase": {
			input: "lower",
			want:  []rune("LOWER"),
		},
		"upper and lower": {
			input: "UPper LowER",
			want:  []rune("UPPER LOWER"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := splitToUppercaseCharacters(tc.input)

			if !slices.Equal(got, tc.want) {
				t.Errorf("got: %c | want: %c", got, tc.want)
			}
		})
	}
}
