package bob

import (
  "regexp"
  "strings"
)

func Hey(remark string) string {
  question := regexp.MustCompile("\\?$")
  letters := regexp.MustCompile("[a-zA-z]+")
  remark = strings.TrimSpace(remark)

  if strings.ToUpper(remark) == remark && question.MatchString(remark) && letters.MatchString(remark) {
    return "Calm down, I know what I'm doing!"
  }

  if question.MatchString(remark) {
    return "Sure."
  }

  if strings.ToUpper(remark) == remark && letters.MatchString(remark) {
    return "Whoa, chill out!"
  }

  if remark == "" {
    return "Fine. Be that way!"
  }

  return "Whatever."
}
