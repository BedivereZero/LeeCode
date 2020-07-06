package algorithms

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	s, i := 0, 0
	for i+nums[i] < len(nums)-1 {
		n, f := i+1, i+1
		for j := i + 1; j <= i+nums[i]; j++ {
			if f < j+nums[j] {
				n, f = j, j+nums[j]
			}
		}
		i = n
		s++
	}
	return s + 1
}
