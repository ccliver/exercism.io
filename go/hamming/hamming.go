package hamming

import (
  "errors"
)

func Distance(a, b string) (int, error) {
  distance := 0

  if len(a) != len(b) {
    return -1, errors.New("Strings not equal length")
  }

  for idx, _ := range a {
    if a[idx] != b[idx] {
      distance += 1
    }
  }

  return distance, nil
}
