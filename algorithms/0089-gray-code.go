package algorithms

func grayCode(n int) []int {
	code := make([]int, 1<<n)
	for i := 0; i < n; i++ {
		for j := 1 << i; j < 1<<(i+1); j++ {
			code[j] = code[(1<<i)-1-(j-(1<<i))] + 1<<i
		}
	}
	return code
}
