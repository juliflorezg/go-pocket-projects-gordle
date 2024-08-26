package gordle

import "strings"

// hint describes the validity of a character in a word.
type hint byte

// feedback is a list of hints, one per character of the word
type feedback []hint

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		{
			return "⬜️"
		}
	case wrongPosition:
		{
			return "🟨"
		}
	case correctPosition:
		{
			return "🟩"
		}
	default:
		{
			// This one should never happen
			return "💔"
		}
	}
}

func (fb feedback) String() string {
	sb := strings.Builder{}

	for _, h := range fb {
		sb.WriteString(h.String())
	}

	return sb.String()
}

// Equal determines equality of two feedbacks.
func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}

	for index, value := range fb {
		if value != other[index] {
			return false
		}
	}

	return true
}
