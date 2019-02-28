package wordcount

import (
	"fmt"
	"regexp"
	"strings"
)

type Frequency map[string]int

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
