package algorithms

import "sort"

func nextPermutation(nums []int)  {
	if len(nums) < 2 {
		return
	}
	previous := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i + 1] {
			previous = i
			break
		}
	}
	if previous > -1 {
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[previous] < nums[i] {
				nums[previous], nums[i] = nums[i], nums[previous]
				break
			}
		}
	}
	sort.Ints(nums[previous + 1:])
}
