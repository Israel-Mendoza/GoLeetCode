package hashmap

import (
	"strconv"
)

// IsHappy isHappy determines if a number is a happy number.
//
// A happy number is defined by the following process:
// 1. Starting with any positive integer, replace the number by the sum of the squares of its digits.
// 2. Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// 3. Those numbers for which this process ends in 1 are happy.
//
// The function returns true if the number is happy, false otherwise.
func IsHappy(n int) bool {
	seenNumbers := make(map[int]bool)
	for n != 1 {
		n = sumOfSquares(numberToDigitSlice(n))
		if _, ok := seenNumbers[n]; ok {
			return false // Cycle detected, not a happy number
		}
		seenNumbers[n] = true
	}
	return true // Reached 1, a happy number
}

// numberToDigitSlice converts a number to a slice of its digits.
func numberToDigitSlice(num int) []int {
	var result []int
	numStr := strconv.Itoa(num)
	for _, char := range numStr {
		integer, _ := strconv.Atoi(string(char))
		result = append(result, integer)
	}
	return result
}

// sumOfSquares calculates the sum of the squares of a slice of numbers.
func sumOfSquares(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number * number
	}
	return sum
}
