package algorithms

func getPivot(nums []int) int {
	h, t := 0, len(nums)
	for h < t {
		if nums[h] < nums[t-1] {
			return t
		}
		if nums[h] == nums[t-1] {
			c := (h + t) / 2
			if nums[h] < nums[c] {
				h = c
			}
			if nums[h] > nums[c] {
				t = c
			}
			if nums[h] == nums[c] {
				for i := c; i > h; i-- {
					if nums[i-1] > nums[i] {
						return i
					}
				}
				for i := c + 1; i < t; i++ {
					if nums[i] < nums[c] {
						return i
					}
				}
				return t
			}
		}
		if nums[h] > nums[t-1] {
			c := (h + t) / 2
			if nums[h] <= nums[c] {
				h = c
			} else {
				t = c
			}
		}
	}
	return h
}

func findMin(nums []int) int {
	pivot := getPivot(nums)
	return nums[pivot%len(nums)]
}
