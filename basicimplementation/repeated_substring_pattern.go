package basicimplementation

import "strings"

// RepeatedSubstringPattern checks if a given string can be formed by repeating a substring.
//
// Parameters:
//
//	s: The input string to check.
//
// Returns:
//
//	true if the string can be formed by repeating a substring, false otherwise.
//
// See my solution on LeetCode: https://leetcode.com/problems/repeated-substring-pattern/solutions/6001035/concatenation-possible-substrings-go
func RepeatedSubstringPattern(s string) bool {
	strLength := len(s)
	for i := 1; i <= strLength/2; i++ {
		if strLength%i == 0 { // Concatenation would have the same length as `s`
			substring := s[:i]
			concatenation := strings.Repeat(substring, strLength/i)
			if concatenation == s {
				return true
			}
		}
	}
	return false
}
