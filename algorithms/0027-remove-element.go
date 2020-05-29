func removeElement(nums []int, val int) int {
	lnums := len(nums)
	switch lnums {
	case 0:
		return 0
	case 1:
		if nums[0] == val {
			return 0
		} else {
			return 1
		}
	}
	l, r := 0, lnums-1
	for l < lnums && nums[l] != val {
		l++
	}
	for 0 <= r && nums[r] == val {
		r--
	}
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		for nums[l] != val {
			l++
		}
		for nums[r] == val {
			r--
		}
	}
	return r + 1
}
