package algorithms

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	for i := 0; i < len(nums)/k; i++ {
		for j := 0; k+i*k+j < len(nums) && j < k; j++ {
			nums[j], nums[k+i*k+j] = nums[k+i*k+j], nums[j]
		}
	}
	rotate(nums[:k], k-len(nums)%k)
}
