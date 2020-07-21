package algorithms

func partition(s string) [][]string {
	results := [][]string{}
	traversal(&results, []string{}, s)
	return results
}

func checkPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func traversal(results *[][]string, current []string, s string) {
	if len(s) == 0 {
		newResult := make([]string, len(current))
		copy(newResult, current)
		*results = append(*results, newResult)
		return
	}

	for i := 0; i < len(s); i++ {
		if checkPalindrome(s[:i+1]) {
			current = append(current, s[:i+1])
			traversal(results, current, s[i+1:])
			current = current[:len(current)-1]
		}
	}
}
