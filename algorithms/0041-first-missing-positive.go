package algorithms

func firstMissingPositive(nums []int) int {
	i := 0
	// create hashmap
	for i < len(nums) {
		if nums[i] == i+1 || nums[i] < 1 || nums[i] > len(nums) || nums[i] == nums[nums[i]-1] {
			i++
		} else {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	// find missing
	for i, n := range nums {
		if i+1 != n {
			return i + 1
		}
	}
	return len(nums) + 1
}
