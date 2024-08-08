package gordle

import (
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
			g := New(strings.NewReader(tc.input))

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v | want =%v ", string(got), string(tc.want))
			}
		})
	}
}
