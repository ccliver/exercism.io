package scrabble

import (
	"regexp"
	"strings"
)

// Score computes the scrabble score of a word.
func Score(word string) int {
	score := 0
	onePoint := regexp.MustCompile("[aeioulnrst]")
	twoPoint := regexp.MustCompile("[dg]")
	threePoint := regexp.MustCompile("[bcmp]")
	fourPoint := regexp.MustCompile("[fhvwy]")
	fivePoint := regexp.MustCompile("[k]")
	eightPoint := regexp.MustCompile("[jx]")
	tenPoint := regexp.MustCompile("[qz]")

	for _, ch := range strings.ToLower(word) {
		switch {
		case onePoint.MatchString(string(ch)):
			score++
		case twoPoint.MatchString(string(ch)):
			score += 2
		case threePoint.MatchString(string(ch)):
			score += 3
		case fourPoint.MatchString(string(ch)):
			score += 4
		case fivePoint.MatchString(string(ch)):
			score += 5
		case eightPoint.MatchString(string(ch)):
			score += 8
		case tenPoint.MatchString(string(ch)):
			score += 10
		}
	}

	return score
}
