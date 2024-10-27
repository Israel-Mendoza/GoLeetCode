package hashmap

import "strings"

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
