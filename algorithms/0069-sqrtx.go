package algorithms

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	a := float64(x) / 2
	for true {
		b := 0.5 * (a + float64(x) / a)
		if int(a) == int(b) {
			break
		}
		a = b
	}
	return int(a)
}
