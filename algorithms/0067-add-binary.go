func reverse(input string) string {
	if len(input) > 0 {
		return reverse(input[1:]) + string([]byte{input[0]})
	} else {
		return input
	}
}

func addBinary(a string, b string) string {
	a = reverse(a)
	b = reverse(b)
	r := ""
	c := 0
	for i :=0; i < len(a) || i < len(b); i++ {
		var x, y int
		if i < len(a) {
			x = int(a[i] - '0')
		} else {
			x = 0
		}
		if i < len(b) {
			y = int(b[i] - '0')
		} else {
			y = 0
		}
		switch x + y + c {
		case 0:
			c, r = 0, r + "0"
		case 1:
			c, r = 0, r + "1"
		case 2:
			c, r = 1, r + "0"
		case 3:
			c, r = 1, r + "1"
		}
	}
	if c > 0 {
		r += "1"
	}
	return reverse(r)
}
