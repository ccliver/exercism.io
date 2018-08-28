package scrabble

import (
	"regexp"
	"strings"
)

// Score computes the scrabble score of a word.
func Score(word string) int {
	score := 0
	points := map[int]*regexp.Regexp{
		1:  regexp.MustCompile("[aeioulnrst]"),
		2:  regexp.MustCompile("[dg]"),
		3:  regexp.MustCompile("[bcmp]"),
		4:  regexp.MustCompile("[fhvwy]"),
		5:  regexp.MustCompile("[k]"),
		8:  regexp.MustCompile("[jx]"),
		10: regexp.MustCompile("[qz]"),
	}

	for pointValue, letters := range points {
		if count := letters.FindAllString(strings.ToLower(word), -1); len(count) > 0 {
			score += len(count) * pointValue
		}
	}

	return score
}
