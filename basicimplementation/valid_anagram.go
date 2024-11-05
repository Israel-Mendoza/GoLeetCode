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
// See my solution on LeetCode: https://leetcode.com/problems/valid-anagram/solutions/6009380/beats-100-in-go-time-o-n-space-o-1.
func IsAnagram(s string, t string) bool {
	// Create an array to store the frequency of each letter
	charFrequency := [26]int{}

	// Count the frequency of each character in string s
	for _, char := range s {
		charFrequency[charPositionInAlphabet(char)]++
	}

	// Decrement the frequency for each character in string t
	// If any frequency becomes negative, t has more of that character than s.
	for _, char := range t {
		letterPosition := charPositionInAlphabet(char)
		charFrequency[letterPosition]--
		if charFrequency[letterPosition] < 0 {
			return false
		}
	}
	// If we made it through both strings without returning false,
	// the strings are anagrams.
	return true
}

// charPositionInAlphabet calculates the zero-based position of a lowercase letter in the alphabet.
// For example, 'a' returns 0, 'b' returns 1, etc.
// This function assumes that the input is a lowercase English letter.
func charPositionInAlphabet(char int32) int {
	// Subtract the ASCII value of 'a' (97) to get the position
	// 'a' — 'a' = 0, 'b' — 'a' = 1, etc.
	return int(char) - 97
}
