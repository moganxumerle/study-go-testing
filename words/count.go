package words

import "strings"

func Count(s, wordFilter string) int {
	wordsArray := strings.Split(s, " ")
	words := make([]string, 0)

	for _, word := range wordsArray {
		if word == "" {
			continue
		}
		words = append(words, word)
	}

	if wordFilter != "" {
		return filter(words, wordFilter)
	}

	return len(words)
}

func filter(words []string, w string) (total int) {
	for _, word := range words {
		if word == w {
			total++
		}
	}
	return
}
