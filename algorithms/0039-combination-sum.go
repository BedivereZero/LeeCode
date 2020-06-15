package algorithms

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var results [][]int
	search(&results, candidates, target, []int {})
	return results
}

func search(resultsPoint *[][]int, candidates []int, target int, current []int){
	if target == 0 {
		new := make([]int, len(current))
		copy(new, current)
		*resultsPoint = append(*resultsPoint, new)
		return
	}
	if target < 0 {
		return
	}
	for point, each := range candidates {
		search(resultsPoint, candidates[point:], target - each, append(current, each))
	}
}
