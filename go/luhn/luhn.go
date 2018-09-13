package luhn

import (
  "regexp"
  "strings"
  "strconv"
)

// Valid determines whether a number is valid per the Luhn formula.
func Valid(number string) bool {
  nonDigits := regexp.MustCompile(`[\D]+`)
  number = strings.Replace(number, " ", "", -1)

  if nonDigits.MatchString(number) || len(number) == 1 {
    return false
  }

  digits := make([]int, len(number))
  secondDigit := false
  for i := len(number) - 1; i >= 0; i-- {
    digit, _ := strconv.Atoi(string(number[i]))

    if secondDigit {
      digits[i] = digit * 2
      if digits[i] > 9 {
        digits[i] -= 9
      }
      secondDigit = false
    } else {
      digits[i] = digit
      secondDigit = true
    }
  }

  sum := 0
  for _, digit := range digits {
    sum += digit
  }

  return sum % 10 == 0
}
