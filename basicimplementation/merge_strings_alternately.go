package basicimplementation

import "strings"

// MergeAlternately Creates a new string builder to efficiently construct the merged string.
//
// See my solution at https://leetcode.com/problems/merge-strings-alternately/solutions/5995269/efficient-string-builder-approach-in-go
func MergeAlternately(word1 string, word2 string) string {
	result := strings.Builder{}

	var majorLength int

	if len(word1) > len(word2) {
		majorLength = len(word1)
	} else {
		majorLength = len(word2)
	}

	for i := 0; i < majorLength; i++ {
		if i < len(word1) {
			result.WriteString(string(word1[i]))
		}

		if i < len(word2) {
			result.WriteString(string(word2[i]))
		}
	}

	return result.String()
}
