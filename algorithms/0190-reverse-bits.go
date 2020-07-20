package algorithms

func reverseBits(num uint32) uint32 {
	var t uint32
	for num > 0 {
		t = (t << 1) + num%2
		num = num >> 1
	}
	return t
}
