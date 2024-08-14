package gordle

import "testing"

func TestFeedbackString(t *testing.T) {
	tt := map[string]struct {
		input feedback
		want  string
	}{
		"absent state": {
			input: feedback{absentCharacter},
			want:  "⬜️",
		},
		"wrong position state": {
			input: feedback{wrongPosition},
			want:  "🟨",
		},
		"correct position state": {
			input: feedback{correctPosition},
			want:  "🟩",
		},
		"two states": {
			input: feedback{absentCharacter, wrongPosition},
			want:  "⬜️🟨",
		},
		"all three states": {
			input: feedback{absentCharacter, wrongPosition, correctPosition},
			want:  "⬜️🟨🟩",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.input.String()

			if got != tc.want {
				t.Errorf("got: %s, but want: %s", got, tc.want)
			}
		})
	}
}
