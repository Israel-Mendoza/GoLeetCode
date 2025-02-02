package easy

func MoveZeroes(nums []int) {
	currentIndex := 0

	for index, number := range nums {
		if number != 0 {
			nums[index], nums[currentIndex] = nums[currentIndex], nums[index]
			currentIndex++
		}
	}
}
