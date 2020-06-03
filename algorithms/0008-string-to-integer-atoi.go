package algorithms

func myAtoi(str string) int {
	int_min, int_max := -1 << 31, 1 << 31 - 1
	h := -1
	for i, c := range str {
		if c != ' ' {
			h = i
			break
		}
	}
	if h < 0 {
		return 0
	}

	if (str[h] < '0' || '9' < str[h]) && str[h] != '+' && str[h] != '-' {
		return 0
	}

	var coefficient int
	switch str[h] {
	case '+':
		coefficient = 1
		h++
	case '-':
		coefficient = -1
		h++
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		coefficient = 1
	}

	t := -1
	for i := h; i < len(str); i++ {
		if str[i] < '0' || '9' < str[i] {
			t = i
			break
		}
	}
	if t < 0 {
		t = len(str)
	}
	if h == t {
		return 0
	}

	// fmt.Println("h:", h, "t:", t)

	integer := coefficient * int(str[h] - '0')

	for _, c := range str[h+1:t] {
		i := int(c - '0')
		if coefficient < 0 {
			if integer < int_min / 10 {
				return int_min
			} else {
				integer = integer * 10
			}
			if integer < int_min + i {
				return int_min
			} else {
				integer = integer - i
			}
		} else {
			if integer > int_max / 10 {
				return int_max
			} else {
				integer = integer * 10
			}
			if integer > int_max - i {
				return int_max
			} else {
				integer = integer + i
			}
		}
	}

	return integer
}
