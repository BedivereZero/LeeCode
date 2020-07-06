package algorithms

func jump(nums []int) int {
	step, end := 0, len(nums)-1
	for end > 0 {
		for i := 0; i < end; i++ {
			if i+nums[i] >= end {
				end = i
				break
			}
		}
		step++
	}
	return step
}
