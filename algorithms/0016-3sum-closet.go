func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closet_distance := -1
	var closet_target int
	for a := range nums {
		if a > 0 && nums[a] == nums[a - 1] {
			continue
		}
		b, c := a + 1, len(nums) - 1
		for b < c {
			if b > (a + 1) && c < len(nums) - 1 && nums[b] == nums[b - 1] && nums[c] == nums[c + 1] {
				b++
				c--
				continue
			}
			s := nums[a] + nums[b] + nums[c]
			var distance int
			if s < target {
				distance = target - s
				b++
			} else {
				distance = s - target
				c--
			}
			if closet_distance < 0 || distance < closet_distance {
				closet_distance, closet_target = distance, s
			}
			if distance == 0 {
				return target
			}
		}
	}
	return closet_target
}
