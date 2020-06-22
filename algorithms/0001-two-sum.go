package algorithms

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, e := m[target -nums[i]]; e {
			return []int {j, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}
