package algorithms

func findPeakElement(nums []int) int {
	n, _ := findPeak(nums, 0, len(nums))
	return n
}

func findPeak(nums []int, head int, tail int) (int, bool) {
	switch tail - head {
	case 1:
		if (head == 0 || nums[head] > nums[head-1]) && (tail == len(nums) || nums[tail-1] > nums[tail]) {
			return head, true
		} else {
			return 0, false
		}
	case 2:
		if nums[head] > nums[tail-1] && (head == 0 || nums[head] > nums[head-1]) {
			return head, true
		}
		if nums[head] < nums[tail-1] && (tail == len(nums) || nums[tail-1] > nums[tail]) {
			return tail - 1, true
		}
		return 0, false
	}

	mid := (head + tail) / 2
	if peak, ok := findPeak(nums, head, mid); ok {
		return peak, ok
	}
	return findPeak(nums, mid, tail)
}
