package algorithms

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	pivot := getPivot(nums)
	var index int
	if target < nums[0] {
		index = binarySearch(nums[pivot:], target) + pivot
	} else {
		index = binarySearch(nums[:pivot], target)
	}
	return index < len(nums) && nums[index] == target
}

func getPivot(nums []int) int {
	h, t := 0, len(nums)
	for h < t {
		if nums[h] < nums[t - 1] {
			h = t
			continue
		}
		if nums[h] > nums[t - 1] {
			m := (h + t) / 2
			if nums[h] > nums[m] {
				t = m
			} else {
				h = m
			}
			continue
		}
		if nums[h] == nums[t - 1] {
			m := (h + t) / 2
			if nums[h] > nums[m] {
				t = m
			} else if nums[h] < nums[m] {
				h = m
			} else {
				h++
			}
			continue
		}
	}
	return h
}

func binarySearch(nums []int, target int) int{
	h, t := 0, len(nums)
	for h < t {
		c := (h + t) / 2
		if nums[c] == target {
			return c
		} else if nums[c] < target {
			h = c + 1
		} else {
			t = c
		}
	}
	return t
}
