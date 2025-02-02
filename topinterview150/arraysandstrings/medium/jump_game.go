package medium

func CanJump(nums []int) bool {
	currentGoal := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		availableJumps := nums[i]
		if i+availableJumps >= currentGoal {
			currentGoal = i
		}
	}

	return currentGoal == 0
}
