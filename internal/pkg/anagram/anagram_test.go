package anagram

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordToSortedLowercaseRunes(t *testing.T) {
	tt := []struct {
		word  string
		runes []rune
	}{
		{
			word:  "foobar",
			runes: []rune{'a', 'b', 'f', 'o', 'o', 'r'},
		},
		{
			word:  "barfoo",
			runes: []rune{'a', 'b', 'f', 'o', 'o', 'r'},
		},
		{
			word:  "boofar",
			runes: []rune{'a', 'b', 'f', 'o', 'o', 'r'},
		},
		{
			word:  "живу",
			runes: []rune{'в', 'ж', 'и', 'у'},
		},
		{
			word:  "вижу",
			runes: []rune{'в', 'ж', 'и', 'у'},
		},
		{
			word:  "Abba",
			runes: []rune{'a', 'a', 'b', 'b'},
		},
		{
			word:  "BaBa",
			runes: []rune{'a', 'a', 'b', 'b'},
		},
	}

	for i, test := range tt {
		t.Run(fmt.Sprintf("step %d", i+1), func(t *testing.T) {
			expected := test.runes
			actual := wordToSortedLowercaseRunes(test.word)
			require.ElementsMatch(t, expected, actual)
		})
	}
}

func Test_BaseAnagram(t *testing.T) {
	tt := []struct {
		word     string
		baseWord string
	}{
		{
			word:     "foobar",
			baseWord: "abfoor",
		}, {
			word:     "foobar",
			baseWord: "abfoor",
		}, {
			word:     "barfoo",
			baseWord: "abfoor",
		}, {
			word:     "живу",
			baseWord: "вжиу",
		}, {
			word:     "Abba",
			baseWord: "aabb",
		}, {
			word:     "abba",
			baseWord: "aabb",
		},
	}

	for i, test := range tt {
		t.Run(fmt.Sprintf("step %d", i+1), func(t *testing.T) {
			expected := test.baseWord
			actual := BaseAnagram(test.word)
			require.Equal(t, expected, actual)
		})
	}
}

func Test_IsAnagram(t *testing.T) {
	tt := []struct {
		wordFirst  string
		wordSecond string
		result     bool
	}{
		{
			wordFirst:  "foobar",
			wordSecond: "barfoo",
			result:     true,
		}, {
			wordFirst:  "foobar",
			wordSecond: "boofar",
			result:     true,
		}, {
			wordFirst:  "barfoo",
			wordSecond: "boofar",
			result:     true,
		}, {
			wordFirst:  "живу",
			wordSecond: "вижу",
			result:     true,
		}, {
			wordFirst:  "Abba",
			wordSecond: "BaBa",
			result:     true,
		}, {
			wordFirst:  "abba",
			wordSecond: "bba",
			result:     false,
		},
	}

	for i, test := range tt {
		t.Run(fmt.Sprintf("step %d", i+1), func(t *testing.T) {
			expected := test.result
			actual := IsAnagram(test.wordFirst, test.wordSecond)
			require.Equal(t, expected, actual)
		})
	}
}
