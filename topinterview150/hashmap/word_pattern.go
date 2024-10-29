package hashmap

import "strings"

// WordPattern determines if the given pattern matches the given string.
//
// A pattern string consists of only lowercase English letters, and a string s of lowercase English words.
// A pattern matches a string s if we can substitute each letter in the pattern with a unique word from s such that the substituted words are the same order as s and no letter is substituted with two different words.
//
// Args:
//
//	pattern: The pattern string.
//	s: The string to match against the pattern.
//
// Returns:
//
//	true if the pattern matches the string, false otherwise.
func WordPattern(pattern string, s string) bool {
	sArray := strings.Split(s, " ")

	if len(sArray) != len(pattern) {
		return false
	}

	mapping := make(map[byte]string)
	seen := make(map[string]bool)

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := sArray[i]

		if existingWord, ok := mapping[char]; ok {
			if existingWord != word {
				return false
			}
		} else {
			if seen[word] {
				return false
			}
			mapping[char] = word
			seen[word] = true
		}
	}

	return true
}
