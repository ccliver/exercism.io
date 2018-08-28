package etl

import "strings"

// Transform converts from a legacy scrabble score to a new version.
func Transform(scores map[int][]string) map[string]int {
  transformedScores := make(map[string]int)

  for score, letters := range scores {
    for _, letter := range letters {
      transformedScores[strings.ToLower(letter)] = int(score)
    }
  }

  return transformedScores
}
