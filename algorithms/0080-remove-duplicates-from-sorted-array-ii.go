package algorithms

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	a, b, c  := 0, 0, 0
	for c < len(nums) || a < b  {
		if c == len(nums) {
			a = b
		} else if a + 2 == b || (a < b && nums[a] != nums[c]) {
			for c < len(nums) && nums[a] == nums[c] {
				c++
			}
			a = b
		} else {
			nums[b], nums[c] = nums[c], nums[b]
			b++
			c++
		}
	}
	return a
}
