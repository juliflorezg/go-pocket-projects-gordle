package gordle

import "testing"

func TestReadCorpus(t *testing.T) {
	tt := map[string]struct {
		file   string
		length int
		err    error
	}{
		"English corpus": {
			file:   "../corpus/english.txt",
			length: 34,
			err:    nil,
		},
		"Empty corpus": {
			file:   "../corpus/empty.txt",
			length: 0,
			err:    ErrEmptyCorpus,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			words, err := ReadCorpus(tc.file)

			if tc.err != err {
				t.Errorf("expected err %v, got %v", tc.err, err)
			}

			if tc.length != len(words) {
				t.Errorf("wrong length, expected %d, got %d", tc.length, len(words))
			}
		})
	}
}
