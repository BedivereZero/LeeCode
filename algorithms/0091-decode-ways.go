package algorithms

func numDecodings(s string) int {
	if len(s) == 1 {
		return waysOfOne(s[0])
	}
	a, b := waysOfOne(s[0]), waysOfOne(s[0]) * waysOfOne(s[1]) + waysOfTwo(s[0], s[1])
	for i := 2; i < len(s); i++ {
		a, b = b, waysOfOne(s[i]) * b + waysOfTwo(s[i - 1], s[i]) * a
	}
	return b
}

func waysOfOne(a byte) int {
	if a > '0' {
		return 1
	}
	return 0
}

func waysOfTwo(a, b byte) int {
	n := int(a - '0') * 10 + int(b - '0')
	if 9 < n && n < 27 {
		return 1
	}
	return 0
}
