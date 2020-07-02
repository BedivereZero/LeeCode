package algorithms

func longestValidParentheses(s string) int {
	longest := 0
	if len(s) < 3 {
		if s == "()" {
			return 2
		} else {
			return 0
		}
	}
	f := make([]int, len(s))
	if s[:2] == "()" {
		f[0], f[1] = 0, 2
		longest = 2
	} else {
		f[0], f[1] = 0, 0
	}
	for i := 2; i < len(f); i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			f[i] = f[i-2] + 2
		} else if i-1-f[i-1] >= 0 && s[i-1-f[i-1]] == '(' {
			if i-2-f[i-1] >= 0 {
				f[i] = f[i-2-f[i-1]] + f[i-1] + 2
			} else {
				f[i] = f[i-1] + 2
			}
		}
		if longest < f[i] {
			longest = f[i]
		}
	}
	return longest
}
