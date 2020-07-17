package algorithms

func singleNumber(nums []int) int {
	h, l := 0, 0
	for i := 0; i < len(nums); i++ {
		h, l = (^nums[i]&h)|(nums[i]&l), (^nums[i]&l)|(nums[i] & ^(h^l))
	}
	return l
}
