package isogram

import "strings"

// IsIsogram tests whether the word is an isogram.
func IsIsogram(word string) bool {
  charactersSeen := make(map[rune]int)
  const seen = 1

  for _, character := range strings.ToLower(word) {
    if len(word) == 0 || character == ' ' || character == '-' {
      continue
    }

    if _, ok := charactersSeen[character]; ok {
      return false
    }
    charactersSeen[character] = seen
  }

  return true
}
