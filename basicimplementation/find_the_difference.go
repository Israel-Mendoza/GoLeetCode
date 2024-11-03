package basicimplementation

// FindTheDifference finds the single character that occurs once in t that does not occur in s.
// See my solution: https://leetcode.com/problems/find-the-difference/solutions/5999361/frequency-based-approach-with-a-map-go
func FindTheDifference(s string, t string) byte {
	// Create a map to store the frequency of each character in s
	availableLetters := make(map[byte]int)
	// Populate the map with character frequencies from s
	populateByteIntMap(availableLetters, s)
	// Find the character in t that has a frequency difference
	return getMissingByte(availableLetters, t)
}

// populateByteIntMap populates a map with the frequency of each character in a given string.
func populateByteIntMap(strIntMap map[byte]int, s string) {
	for _, char := range s {
		letter := byte(char)
		strIntMap[letter]++
	}
}

// getMissingByte finds the single character in s that does not occur in strIntMap.
func getMissingByte(strIntMap map[byte]int, s string) byte {
	for _, char := range s {
		letter := byte(char)
		if strIntMap[letter] == 0 {
			return letter
		}
		strIntMap[letter]--

		// Removing entry if count is zero
		if strIntMap[letter] == 0 {
			delete(strIntMap, letter)
		}
	}
	return 0 // If no difference is found, return 0
}
