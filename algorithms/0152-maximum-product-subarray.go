package algorithms

func maxProduct(nums []int) int {
	maxs := make([]int, len(nums))
	mins := make([]int, len(nums))

	maximum := nums[0]
	maxs[0] = nums[0]
	mins[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i] < 0:
			mins[i] = min(nums[i], maxs[i-1]*nums[i])
			maxs[i] = max(nums[i], mins[i-1]*nums[i])
		case nums[i] == 0:
			mins[i] = 0
			maxs[i] = 0
		case nums[i] > 0:
			mins[i] = min(nums[i], mins[i-1]*nums[i])
			maxs[i] = max(nums[i], maxs[i-1]*nums[i])
		}
		if maxs[i] > maximum {
			maximum = maxs[i]
		}
		maximum = max(maximum)
	}
	return maximum
}
