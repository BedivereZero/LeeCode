package algorithms

func isValid(s string) bool {
	if s == "" {
		return true
	}
	a := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]byte, 0)
	for _, i := range s {
		j := byte(i)
		if j == '(' || j == '[' || j == '{' {
			stack = append(stack, j)
		} else if len(stack) > 0 && a[stack[len(stack)-1]] == j {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
