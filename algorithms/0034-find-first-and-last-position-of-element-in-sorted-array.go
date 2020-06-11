package algorithms

func searchRange(nums []int, target int) []int {
	result := []int {-1, -1}
	if len(nums) < 1 {
		return result
	}
	h, t := 0, len(nums)
	for h < t {
		c := (h + t) / 2
		if nums[c] < target {
			h = c + 1
		} else {
			t = c
		}
	}
	if h == len(nums) || nums[h] != target {
		return result
	}
	result[0] = h
	h, t = 0, len(nums)
	for h < t {
		c := (h + t) / 2
		if nums[c] <= target {
			h = c + 1
		} else {
			t = c
		}
	}
	result[1] = h - 1
	return result
}
