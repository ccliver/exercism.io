package hamming

import (
  "errors"
)

// Distance determines the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
  distance := 0

  if len(a) != len(b) {
    return -1, errors.New("Strings not equal length")
  }

  for idx := range a {
    if a[idx] != b[idx] {
      distance++
    }
  }

  return distance, nil
}
