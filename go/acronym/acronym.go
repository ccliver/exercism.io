package acronym

import (
  "strings"
  "regexp"
)

func Abbreviate(s string) string {
  acronym := ""
  separators := regexp.MustCompile("[ -]")

  for _, word := range separators.Split(s, -1) {
    acronym += strings.ToUpper(string(word[0]))
  }

	return acronym
}
