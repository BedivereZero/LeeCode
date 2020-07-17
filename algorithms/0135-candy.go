package algorithms

func candy(ratings []int) int {
	cache := make([]int, len(ratings))
	cache[0] = 0
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			cache[i] = cache[i-1] + 1
		}
	}
	p := 0
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && cache[i] < p+1 {
			cache[i] = p + 1
		}
		p = cache[i]
	}
	sum := 0
	for _, c := range cache {
		sum += c
	}
	return sum + len(ratings)
}
