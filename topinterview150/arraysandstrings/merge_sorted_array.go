package arraysandstrings

import "math"

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	idx1 := m - 1
	idx2 := n - 1

	for currentIndex := len(nums1) - 1; currentIndex >= 0; currentIndex-- {
		firstNumber := getFromArray(nums1, idx1)
		secondNumber := getFromArray(nums2, idx2)
		if secondNumber >= firstNumber {
			nums1[currentIndex] = secondNumber
			idx2--
		} else {
			nums1[currentIndex] = firstNumber
			idx1--
		}
	}

}

func getFromArray(array []int, index int) int {
	if index < 0 {
		return math.MinInt
	}
	return array[index]
}
