package algorithms

func canJump(nums []int) bool {
	result := make([]bool, len(nums))
	result[0] = true
	for i := 0; i < len(nums); i++ {
		if ! result[i] {
			continue
		}
		for j := i + 1; j < len(nums) && j <= i + nums[i]; j++ {
			result[j] = true
		}
	}
	return result[len(nums) - 1]
}
