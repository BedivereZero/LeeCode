package algorithms

func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if 0 < nums[i - 1] {
			nums[i] = nums[i - 1] + nums[i]
		}
	}
	max := nums[0]
	for _, each := range nums[1:] {
		if max < each {
			max = each
		}
	}
	return max
}
