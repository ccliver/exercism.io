package anagram

import (
	"sort"
	"strings"
)

func Detect(word string, candidates []string) (matches []string) {
	for _, nextWord := range candidates {
		if strings.ToLower(word) == strings.ToLower(nextWord) {
			continue
		}

		if checkAnagram(strings.ToLower(word), strings.ToLower(nextWord)) {
			matches = append(matches, nextWord)
		}
	}
	return
}

func checkAnagram(a string, b string) bool {
	word := strings.Split(a, "")
	sort.Strings(word)
	candidate := strings.Split(b, "")
	sort.Strings(candidate)

	if strings.Compare(strings.Join(word, ""), strings.Join(candidate, "")) == 0 {
		return true
	}

	return false
}
