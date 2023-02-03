package words_test

import (
	"testing"

	"github.com/moganxumerle/study-go-testing/words"
)

func TestCount(t *testing.T) {
	testCases := []struct {
		Text     string
		Filter   string
		Expected int
	}{
		{
			Text:     "texto de teste",
			Expected: 3,
		},
		{
			Text:     "texto de teste de 6 palavras",
			Expected: 6,
		},
		{
			Text:     "texto de teste de 2 palavras filtradas",
			Filter:   "de",
			Expected: 2,
		},
		{
			Text:     "texto com 2  espaços",
			Expected: 4,
		},
		{
			Text:     "",
			Expected: 0,
		},
		{
			Text:     "texto de teste de 2 palavras filtradas",
			Filter:   "letra",
			Expected: 0,
		},
	}

	for _, tc := range testCases {
		c := words.Count(tc.Text, tc.Filter)
		if c != tc.Expected {
			t.Fatalf("expected: %d, got: %d", tc.Expected, c)
		}
	}
}
