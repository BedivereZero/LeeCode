package algorithms

func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		if nums[0] < nums[1] {
			return nums[1]
		} else {
			return nums[0]
		}
	}

	if nums[0] > nums[1] {
		nums[1] = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		nums[i] += nums[i-2]
		if nums[i] < nums[i-1] {
			nums[i] = nums[i-1]
		}
	}
	return nums[len(nums)-1]
}
