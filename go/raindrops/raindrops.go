package raindrops

import "strconv"

// Convert a number to a raindrop string.
func Convert(n int) string{
  convertedPhrase := ""

  if n % 3 == 0 {
    convertedPhrase += "Pling"
  }
  if n % 5 == 0 {
    convertedPhrase += "Plang"
  }
  if n % 7 == 0 {
    convertedPhrase += "Plong"
  }

  if convertedPhrase == "" {
    return strconv.Itoa(n)
  }

  return convertedPhrase
}
