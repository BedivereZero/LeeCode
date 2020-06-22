package algorithms

func reverse(input []int) []int {
	if len(input) == 0 {
		return input
	} else {
		return(append(reverse(input[1:]), input[0]))
	}
}

func plusOne(digits []int) []int {
	digits = reverse(digits)
	c := 1
	for i := 0; i < len(digits); i++ {
		digits[i], c = (digits[i] + c) % 10, (digits[i] + c) / 10
	}
	if c > 0 {
		digits = append(digits, c)
	}
	return reverse(digits)
}
