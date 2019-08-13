package anagram

import (
	"sort"
	"strings"
)

func wordToSortedLowercaseRunes(word string) []rune {
	wordLowercase := strings.ToLower(word)
	wr := make([]rune, 0, len(wordLowercase))
	for _, l := range wordLowercase {
		wr = append(wr, l)
	}

	sort.Slice(wr, func(i int, j int) bool {
		return wr[i] < wr[j]
	})
	return wr
}

func BaseAnagram(word string) string {
	runes := wordToSortedLowercaseRunes(word)
	return string(runes)
}

func IsAnagram(lhs string, rhs string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	lhsRunes := wordToSortedLowercaseRunes(lhs)
	rhsRunes := wordToSortedLowercaseRunes(rhs)

	for i := range lhsRunes {
		if lhsRunes[i] != rhsRunes[i] {
			return false
		}
	}
	return true
}
