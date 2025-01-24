package topinterview150

func MajorityElement(nums []int) int {

	element := nums[0]
	count := 0

	for _, num := range nums {
		if num == element {
			count++
			continue
		}

		count--
		if count == 0 {
			element = num
			count++
		}
	}

	return element
}
