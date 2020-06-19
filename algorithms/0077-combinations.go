package algorithms

func combine(n int, k int) [][]int {
	length := combinate(n, k)
	matrix := make([][]int, length)
	current := start(k)
	for i := range matrix {
		matrix[i] = append(matrix[i], current...)
		next(current, n)
	}
	return matrix
}

func combinate(n int, k int) int {
	if k > n / 2 {
		k = n - k
	}
	base := 1
	for i := 0; i < k; i++ {
		base = base * (n - i) / (i + 1)
	}
	return base
}

func start(n int) []int {
	array := make([]int, n)
	for i := range array {
		array[i] = i + 1
	}
	return array
}

func next(current []int, maximum int) {
	var position int
	for i := len(current); i > 0; i-- {
		if current[i - 1] < maximum - len(current) + i {
			position = i - 1
			break
		}
	}
	current[position]++
	for i := position + 1; i < len(current); i++ {
		current[i] = current[i - 1] + 1
	}
}
