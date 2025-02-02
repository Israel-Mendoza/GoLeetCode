package easy

func RemoveElement(nums []int, val int) int {
	availableIndex := 0

	for _, num := range nums {
		if num != val {
			nums[availableIndex] = num
			availableIndex++
		}
	}

	return availableIndex
}
