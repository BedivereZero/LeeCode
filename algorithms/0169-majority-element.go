package algorithms

func majorityElement(nums []int) int {
	record := map[int]int{}
	for _, n := range nums {
		if v, _ := record[n]; v+1 > len(nums)/2 {
			return n
		}
		record[n]++
	}
	return 0
}
