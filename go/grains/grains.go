package grains

import (
  "errors"
)

const numSquares = 64

// Square returns the number of grains on a square of the chessboard.
func Square(x int) (uint64, error) {
  if (x < 1 || x > numSquares) {
    return uint64(x), errors.New("Invalid input")
  }

  if (x == 1) {
    return uint64(x), nil
  }

  return 2 << (uint64(x) - 2), nil
}

// Total returns the total number of grains on the chessboard.
func Total() uint64 {
  return 2 << (uint64(numSquares) - 1) - 1
}
