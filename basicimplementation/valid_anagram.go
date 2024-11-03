package basicimplementation

// IsAnagram determines if two strings are anagrams of each other.
//
// An anagram is a word or phrase formed by rearranging the letters of another word or phrase.
//
// Parameters:
//
//	s: The first string to compare.
//	t: The second string to compare.
//
// Returns:
//
//	true if `s` and `t` are anagrams, false otherwise.
//
// See my solution on LeetCode: https://leetcode.com/problems/valid-anagram/solutions/6000641/using-frequency-map-go
func IsAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sChars := getStringCharFrequency(s)

	for _, char := range t {
		letter := byte(char)
		if sChars[letter] <= 0 {
			return false
		}
		decreaseCharFrequency(sChars, letter)
	}

	return true
}

// decreaseCharFrequency decrements the frequency of a character in a given character frequency map.
//
// Parameters:
//
//	charMap: The character frequency map.
//	char: The character to decrement.
func decreaseCharFrequency(charMap map[byte]int, char byte) {
	if count, ok := charMap[char]; ok {
		if count > 1 {
			charMap[char]--
		} else {
			delete(charMap, char)
		}
	}
}

// getStringCharFrequency calculates the frequency of each character in a given string.
//
// Parameters:
//
//	s: The input string.
//
// Returns:
//
//	A map containing the frequency of each character in the string.
func getStringCharFrequency(s string) map[byte]int {
	availableChars := make(map[byte]int)
	for _, char := range s {
		letter := byte(char)
		availableChars[letter]++
	}
	return availableChars
}
