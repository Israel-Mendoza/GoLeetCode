package basicimplementation

// StrStr implements a naive string searching algorithm to find the first occurrence of the needle string within the haystack string.
//
// Args:
//
//	haystack: The string to be searched.
//	needle:   The string to be found.
//
// Returns:
//
//	The index of the first occurrence of the needle in the haystack, or -1 if not found.
//
// See my solution on LeetCode: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/solutions/6000456/optimized-substring-search-go
func StrStr(haystack string, needle string) int {
	switch {
	case len(haystack) == 0: // haystack is empty
		return -1
	case len(needle) > len(haystack): // needle can't be longer than haystack
		return -1
	case len(haystack) == len(needle) && haystack == needle: // haystack and needle are equal
		return 0
	}

	// Looking for the index of the needle in the haystack
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		letter := haystack[i]
		if letter == needle[0] {
			substring := haystack[i : i+len(needle)]
			if needle == substring {
				return i
			}
		}
	}
	return -1
}
