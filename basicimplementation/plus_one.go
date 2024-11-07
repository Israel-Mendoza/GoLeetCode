package basicimplementation

// PlusOne increments an integer represented as a digit array by one.
//
// The input `digits` is a slice of integers, where each digit[i] represents
// the i-th digit of the integer. The digits are ordered from most significant
// to least significant.
//
// The function modifies the input `digits` slice in-place and returns it.
// If a carry-over occurs to the most significant digit, a new digit 1 is
// appended to the beginning of the slice.
//
// See my solution on LeetCode: https://leetcode.com/problems/plus-one/solutions/6017826/simple-solution-time-o-n-space-o-1-go
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
