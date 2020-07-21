package algorithms

func hammingWeight(num uint32) int {
	w := 0
	for num > 0 {
		if num%2 == 1 {
			w++
		}
		num = num >> 1
	}
	return w
}
