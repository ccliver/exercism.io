package wordcount

import (
	"regexp"
	"strings"
)

// Frequency records the number of times a word appears in a string.
type Frequency map[string]int

// WordCount counts the number of times a word appears in a string and returns a Frequency object containing the counts.
func WordCount(in string) Frequency {
	results := Frequency{}

	for _, word := range regexp.MustCompile(`[,\s]+`).Split(in, -1) {
		word = regexp.MustCompile(`[!&@$%^:.]+`).ReplaceAllString(word, "")
		word = regexp.MustCompile(`'$`).ReplaceAllString(word, "")
		word = regexp.MustCompile(`^'`).ReplaceAllString(word, "")
		word = strings.ToLower(word)
		if !regexp.MustCompile(`['a-z\d]+`).MatchString(word) { // Filter out blanks.
			continue
		}

		results[word]++
	}

	return results
}
