package algorithms

func trailingZeroes(n int) int {
	count2 := countBase(n, 2)
	count5 := countBase(n, 5)
	if count2 < count5 {
		return count2
	}
	return count5
}

func countBase(x, b int) int {
	c := 0
	t := x / b
	for t > 0 {
		c += t
		t = t / b
	}
	return c
}
