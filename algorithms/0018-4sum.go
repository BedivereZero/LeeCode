package algorithms

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for a := 0; a < len(nums) - 3; a++ {
		if a > 0 && nums[a] == nums[a - 1] {
			continue
		}
		for b := a + 1; b < len(nums) - 2; b++ {
			if b > a + 1 && nums[b] == nums[b - 1] {
				continue
			}
			c, d := b + 1, len(nums) - 1
			for c < d {
				if c > b + 1 && d < len(nums) - 1 && nums[c] == nums[c - 1] && nums[d] == nums[d + 1] {
					c++
					d--
					continue
				}
				if nums[a] + nums[b] + nums[c] + nums[d] < target {
					c++
				} else if nums[a] + nums[b] + nums[c] + nums[d] > target {
					d--
				} else {
					result = append(result, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
					d--
				}
			}
		}
	}
	return result
}
