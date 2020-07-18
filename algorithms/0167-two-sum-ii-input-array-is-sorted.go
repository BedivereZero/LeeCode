package algorithms

func twoSum(numbers []int, target int) []int {
	cache := map[int]int{}
	var a, b int
	for i, n := range numbers {
		if v, ok := cache[n]; ok {
			a, b = i, v
		} else {
			cache[target-n] = i
		}
	}
	if a > b {
		a, b = b, a
	}
	return []int{a + 1, b + 1}
}
