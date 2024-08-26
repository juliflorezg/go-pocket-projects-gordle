package gordle

import "testing"

func inCorpus(corpus []string, word string) bool {
	for _, corpusWord := range corpus {
		if word == corpusWord {
			return true
		}
	}
	return false
}

func TestPickWord(t *testing.T) {
	corpus := []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"}
	word := pickWord(corpus)

	if !inCorpus(corpus, word) {
		t.Fatalf("expected a word in the corpus, got %q", word)
	}
}
