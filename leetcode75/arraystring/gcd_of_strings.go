package arraystring

// GcdOfStrings finds the greatest common divisor (GCD) of two strings.
//
// A GCD of strings is the longest string that is a prefix of both str1 and str2.
//
// Args:
//
//	str1: The first string.
//	str2: The second string.
//
// Returns:
//
//	The GCD string if it exists, otherwise an empty string.
func GcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	strLength := gcd(len(str1), len(str2))
	return str1[:strLength]
}

// gcd calculates the greatest common divisor of two integers using the Euclidean algorithm.
//
// Args:
//
//	a: The first integer.
//	b: The second integer.
//
// Returns:
//
//	The GCD of a and b.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
