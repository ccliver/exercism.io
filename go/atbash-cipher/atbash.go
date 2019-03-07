package atbash

import (
	"regexp"
	"strings"
)

func Atbash(text string) (ct string) {
	charMap := map[string]string {
		"a": "z", "b": "y", "c": "x", "d": "w", "e": "v",
		"f": "u", "g": "t", "h": "s", "i": "r", "j": "q",
		"k": "p", "l": "o", "m": "n", "n": "m", "o": "l",
		"p": "k", "q": "j", "r": "i", "s": "h", "t": "g",
		"u": "f", "v": "e", "w": "d", "x": "c", "y": "b",
		"z": "a",
	}
	re := regexp.MustCompile(`[\W]+`)
	characters := 0

	for _, character := range strings.Split(strings.ToLower(re.ReplaceAllString(text, "")), "") {
		if characters == 5 {
			characters = 0
			ct += " "
		}

		if _, ok := charMap[character]; ok {
			ct += charMap[character]
		} else {
			ct += character
		}

		characters++
	}

	return
}
