package gordle

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
