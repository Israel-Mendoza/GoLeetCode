package topinterview150

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	lastUnique := nums[0]
	availableIndex := 1

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > lastUnique {
			nums[availableIndex] = num
			availableIndex++
			lastUnique = num
		}
	}

	return availableIndex
}
