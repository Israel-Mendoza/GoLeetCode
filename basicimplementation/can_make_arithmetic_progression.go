package basicimplementation

import "sort"

// CanMakeArithmeticProgression checks if the given integer array can be rearranged to form an arithmetic progression.
//
// An arithmetic progression is a sequence of numbers such that the difference between any two consecutive terms is the same.
//
// Args:
//
//	arr: The integer array to check.
//
// Returns:
//
//	true if the array can be rearranged to form an arithmetic progression, false otherwise.
//
// See my solution on LeetCode: https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/solutions/6027970/sorting-the-array-go
func CanMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	prev := arr[0]
	diff := arr[1] - prev

	for i := 1; i < len(arr); i++ {
		currentNum := arr[i]
		if currentNum-prev != diff {
			return false
		}
		prev = currentNum
	}

	return true
}
