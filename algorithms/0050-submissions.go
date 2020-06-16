func myPow(x float64, n int) float64 {
    if n < 0 {
        x, n = 1 / x, -n
	}
	switch n {
	case 0:
		return 1
	case 1:
		return x
	default:
		return myPow(x * x, n / 2) * myPow(x, n % 2)
	}
}
