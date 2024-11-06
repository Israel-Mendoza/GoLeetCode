package basicimplementation

// MoveZeroes moves all zeros to the end of the given integer array `nums` while maintaining the relative order of the non-zero elements.
//
// This function operates in-place, modifying the input array directly.
//
// Args:
//
//	nums: The integer array to be modified.
//
// See my solution on LeetCode: https://leetcode.com/problems/move-zeroes/solutions/6014274/time-o-n-space-o-1-beats-100-go
func MoveZeroes(nums []int) {
	availableSlot := 0
	for nums[availableSlot] != 0 && availableSlot < len(nums)-1 {
		availableSlot++
	}

	for i := availableSlot + 1; i < len(nums); i++ {
		currentNum := nums[i]
		if currentNum == 0 {
			continue
		}
		nums[availableSlot] = currentNum
		nums[i] = 0
		availableSlot++
	}
}
