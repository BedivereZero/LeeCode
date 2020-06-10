package algorithms

import "fmt"

func getRotatePivot(nums []int) int {
	h, t := 0, len(nums)
	for h < t {
		if nums[h] <= nums[t - 1] {
			return t
		}
		c := (h + t) / 2
		if nums[h] > nums[c] {
			t = c
		} else {
			h = c
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

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	pivot := getRotatePivot(nums)
	head, tail := 0, len(nums)
	if target < nums[0] {
		head = pivot
	} else {
		tail = pivot
	}
	index := head + binarySearch(nums[head:tail], target)
	if index < len(nums) && nums[index] == target {
		return index
	} else {
		return -1
	}
}
