package algorithms

func permute(nums []int) [][]int {
	var result [][] int
	var current [] int
	search(&result, &current, nums)
	return result
}

func search(resultPoint *[][]int, current *[]int, nums []int) {
	if len(nums) == 0 {
		tmp := make([]int, len(*current))
		copy(tmp, *current)
		*resultPoint = append(*resultPoint, tmp)
	}
	for index, value := range nums {
		*current = append(*current, value)
		var next []int
		next = append(next, nums[:index]...)
		next = append(next, nums[index + 1:]...)
		search(resultPoint, current, next)
		*current = (*current)[:len(*current) - 1]
	}
}
