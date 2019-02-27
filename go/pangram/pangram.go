package pangram

import (
	"regexp"
	"strings"
)

func IsPangram(sentence string) bool {
	letters_seen := ""
	alphabet_re := regexp.MustCompile(`[a-z]+`)

	if len(sentence) == 0 {
		return false
	}

	for _, letter := range strings.Split(strings.ToLower(sentence), "") {
		if alphabet_re.MatchString(letter) {
			if !strings.Contains(letters_seen, letter) {
				letters_seen += letter
			}
		}
	}

	if len(letters_seen) == 26 {
		return true
	}

	return false
}
