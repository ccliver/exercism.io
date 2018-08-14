package hamming

import (
  "strings"
  "errors"
)

func Distance(a, b string) (int, error) {
  distance := 0
  nucleotidesA := strings.Split(a, "")
  nucleotidesB := strings.Split(b, "")

  if len(a) != len(b) {
    return -1, errors.New("Strings not equal length")
  }

  for idx, _ := range nucleotidesA {
    if nucleotidesA[idx] != nucleotidesB[idx] {
      distance += 1
    }
  }

  return distance, nil
}
