package algorithms

func search(result *[]string, current string, left int, right int, remain int) {
	if remain == 0 && left == right {
		*result = append(*result, current)
	}
	if remain > 0 {
		search(result, current + "(", left + 1, right, remain - 1)
	}
	if right < left {
		search(result, current + ")", left, right + 1, remain)
	}
}

func generateParenthesis(n int) []string {
	var result []string
	search(&result, "", 0, 0, n)
	return result
}
