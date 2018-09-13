package diffsquares

// SquareOfSums returns the square of the sum of the first x natural numbers.
func SquareOfSums(x int) int {
  sum := 0

  for i := 1; i <= x; i++ {
    sum += i
  }

  return sum * sum
}

// SumOfSquares returns the sum of the squares of the first x natural numbers.
func SumOfSquares(x int) int {
  sum := 0

  for i := 1; i <= x; i++ {
    sum += i * i
  }

  return sum
}

// Difference returns the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(x int) int {
  return SquareOfSums(x) - SumOfSquares(x)
}
