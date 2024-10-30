package hashmap

// ContainsNearbyDuplicate checks if there are two distinct indices i and j in the array nums
// such that nums[i] == nums[j] and abs(i - j) <= k.
func ContainsNearbyDuplicate(nums []int, k int) bool {
	// Map to store each value and its last index.
	valueToIndex := make(map[int]int)

	for i, value := range nums {
		// If the current value is already in the map and the difference between the current index
		// and the previous index is less than or equal to k, we found a duplicate within the specified distance.
		if j, ok := valueToIndex[value]; ok && intAbs(i-j) <= k {
			return true
		}
		// Update the last index for the current value.
		valueToIndex[value] = i
	}

	return false
}

// intAbs returns the absolute value of an integer.
func intAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
