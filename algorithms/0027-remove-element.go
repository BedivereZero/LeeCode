package algorithms

func removeElement(nums []int, val int) int {
	lnums := len(nums)
	count := 0
	for i := 0; i < lnums; i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
