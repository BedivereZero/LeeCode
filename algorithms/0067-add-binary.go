package algorithms

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	r := ""
	c := 0
	for i := 0; i < la || i < lb; i++ {
		x, y := 0, 0
		if la - i - 1 >= 0 {
			x = int(a[la - i - 1] - '0')
		}
		if lb - i - 1 >= 0 {
			y = int(b[lb - i - 1] - '0')
		}
		switch x + y + c {
		case 0:
			c, r = 0, "0" + r
		case 1:
			c, r = 0, "1" + r
		case 2:
			c, r = 1, "0" + r
		case 3:
			c, r = 1, "1" + r
		}
	}
	if c > 0 {
		r = "1" + r
	}
	return r
}
