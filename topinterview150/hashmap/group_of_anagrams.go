package hashmap

import (
	"slices"
	"strings"
)

// GroupAnagrams groups anagrams together.
//
// Given an array of strings words, the function groups the anagrams together.
// Anagrams are words that have the same characters but in a different order.
//
// Args:
//
//	words: A slice of strings to group.
//
// Returns:
//
//	A list of groups, where each group contains anagrams.
func GroupAnagrams(words []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)
		if _, ok := anagramGroups[sortedWord]; ok {
			anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
		} else {
			anagramGroups[sortedWord] = []string{word}
		}
	}
	var result [][]string

	for _, value := range anagramGroups {
		result = append(result, value)
	}

	return result
}

// sortString sorts a string alphabetically.
//
// Args:
//
//	s: The string to sort.
//
// Returns:
//
//	The sorted string.
func sortString(s string) string {
	charSlice := strings.Split(s, "")
	slices.Sort(charSlice)
	return strings.Join(charSlice, "")
}
