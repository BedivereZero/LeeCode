package algorithms

func sortColors(nums []int) {
	var a, b, c int

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			a, b = i, i
			break
		}
	}

	for i := len(nums); i > 0; i-- {
		if nums[i-1] != 2 {
			c = i
			break
		}
	}

	// nums[:a] are '0'
	// nums[a:b] are '1'
	// nums[c:] are '2'
	// nums[b:c] are unknown

	for b < c {
		switch nums[b] {
		case 0:
			nums[a], nums[b] = nums[b], nums[a]
			a++
			b++
		case 1:
			b++
		case 2:
			nums[b], nums[c-1] = nums[c-1], nums[b]
			c--
		}
	}
}
