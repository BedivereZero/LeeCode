package algorithms

func findMin(nums []int) int {
	pivot := getPivot(nums)
	if pivot < len(nums) {
		return nums[pivot]
	}
	return nums[0]
}

func getPivot(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	i, j := 0, len(nums)
	for i < j {
		h := int(uint(i+j) >> 1)
		if nums[h] >= nums[0] {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}
