package algorithms

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int {{}}
	}
	matrix := make([][]int, 1 << uint(len(nums)))
	for i, num := range nums {
		n := 1 << uint(i)
		for j := 0; j < n; j++ {
			matrix[n + j] = make([]int, len(matrix[j]) + 1)
			copy(matrix[n + j], matrix[j])
			matrix[n + j][len(matrix[j])] = num
		}
	}
	return matrix
}
