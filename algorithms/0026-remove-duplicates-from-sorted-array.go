func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	p, q := 0, 1
	for q < length {
		if nums[p] != nums[q] {
			p++
			nums[p] = nums[q]
		}
		q++
	}
	return p + 1
}
