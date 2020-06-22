package algorithms

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for a := 0; a < len(nums) - 2; a++ {
		if a > 0 && nums[a] == nums[a - 1] {
			continue
		}
		b, c := a + 1, len(nums) - 1
		for b < c {
			if b > a + 1 && c < len(nums) - 1 && nums[b] == nums[b - 1] && nums[c] == nums[c + 1] {
				b++
				c--
				continue
			}
			if nums[a] + nums[b] + nums[c] < 0 {
				b++
			} else if nums[a] + nums[b] + nums[c] > 0 {
				c--
			} else {
				result = append(result, []int {nums[a], nums[b], nums[c]})
				b++
				c--
			}
		}
	}
	return result
}
