func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y, z := 0, x
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	return y == z
}
