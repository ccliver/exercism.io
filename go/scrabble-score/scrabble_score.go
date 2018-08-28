package scrabble

import (
  "strings"
  "regexp"
)

// Score computes the scrabble score of a word.
func Score(word string) int {
  score := 0

  for _, ch := range strings.ToLower(word) {
    if matched, _ := regexp.MatchString("[aeioulnrst]", string(ch)); matched {
      score++
    }
    if matched, _ := regexp.MatchString("[dg]", string(ch)); matched {
      score += 2
    }
    if matched, _ := regexp.MatchString("[bcmp]", string(ch)); matched {
      score += 3
    }
    if matched, _ := regexp.MatchString("[fhvwy]", string(ch)); matched {
      score += 4
    }
    if matched, _ := regexp.MatchString("[k]", string(ch)); matched {
      score += 5
    }
    if matched, _ := regexp.MatchString("[jx]", string(ch)); matched {
      score += 8
    }
    if matched, _ := regexp.MatchString("[qz]", string(ch)); matched {
      score += 10
    }
  }

  return score
}
