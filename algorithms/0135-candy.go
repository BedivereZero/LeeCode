package algorithms

func candy(ratings []int) int {
	sum := 0
	for i := 1; i < len(ratings); {
		// skip equal
		for ; i < len(ratings) && ratings[i] == ratings[i-1]; i++ {
		}

		// count up
		u := 0
		for ; i+u < len(ratings) && ratings[i+u] > ratings[i+u-1]; u++ {
		}
		i += u

		// count down
		d := 0
		for ; i+d < len(ratings) && ratings[i+d] < ratings[i+d-1]; d++ {
		}
		i += d

		// compare and switch u and d, make sure u <= d
		if d < u {
			u, d = d, u
		}

		// calculate candy
		sum += (1+u-1)*(u-1)/2 + (1+d)*d/2
	}
	return sum + len(ratings)
}
