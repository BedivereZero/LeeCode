func intToRoman(num int) string {
	var r string
	var c int

	c, num = num / 1000, num % 1000
	for i := 0; i < c; i++ {
		r += "M"
	}

	c, num = num / 100, num % 100
	switch c {
	case 9:
		r += "CM"
	case 4:
		r += "CD"
	default:
		if c >= 5 {
			r += "D"
		}
		for i := 0; i < c % 5; i ++ {
			r += "C"
		}
	}

	c, num = num / 10, num % 10
	switch c {
	case 9:
		r += "XC"
	case 4:
		r += "XL"
	default:
		if c >= 5 {
			r += "L"
		}
		for i := 0; i < c % 5; i ++ {
			r += "X"
		}
	}

	c = num
	switch c {
	case 9:
		r += "IX"
	case 4:
		r += "IV"
	default:
		if c >= 5 {
			r += "V"
		}
		for i := 0; i < c % 5; i ++ {
			r += "I"
		}
	}

	return r
}
