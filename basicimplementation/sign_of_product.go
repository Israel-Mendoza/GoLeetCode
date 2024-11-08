package basicimplementation

// ArraySign calculates the sign of the product of all elements in the given array.
//
// It iterates through the array, counting the number of negative elements.
// If there's a zero, the product is zero. Otherwise, if the number of
// negative elements is even, the product is positive; otherwise, it's negative.
//
// See my solution on LeetCode: https://leetcode.com/problems/sign-of-the-product-of-an-array/solutions/6021906/best-than-100-in-go-time-o-n-space-o-1/
func ArraySign(nums []int) int {
	numberOfNegatives := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			numberOfNegatives++
		}
	}
	if numberOfNegatives%2 == 0 {
		return 1
	}
	return -1
}
