package arraysandstrings

func LengthOfLastWord(s string) int {
	foundLetter := false
	wordLength := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if foundLetter {
				break
			}
		} else {
			foundLetter = true
			wordLength++
		}
	}

	return wordLength
}
