package algorithms

func longestPalindrome(s string) string {
	lens := len(s)
	if lens < 2 {
		return s
	}
	status := make([]bool, lens * lens)
	lh, lt := 0, 0
	for l := 0; l < lens; l++ {
		for h := 0; h + l < lens; h++ {
			if l < 2 {
				status[h * lens + h + l] = s[h] == s[h + l]
			} else {
				status[h * lens + h + l] = status[(h + 1) * lens + h + l - 1] && s[h] == s[h + l]
			}
			if status[h * lens + h + l] && lt - lh < l {
				lh, lt = h, h + l
			}
		}
	}
	return s[lh:lt + 1]
}
