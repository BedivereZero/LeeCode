// package algorithms
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int {10, 1, 2, 7, 6, 1, 5}
	t := 8
	for _, v := range combinationSum2(s, t) {
		 fmt.Println(v)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	var results [][]int
	sort.Ints(candidates)
	search(&results, &candidates, 0, []int{}, target)
	return results
}

func search(resultsPoint *[][]int, candidatesPoint *[]int, position int, current []int, target int) {
	if target == 0 {
		new := make([]int, len(current))
		copy(new, current)
		*resultsPoint = append(*resultsPoint, new)
		return
	} else if target < 0 {
		return
	}
	for i := position; i < len(*candidatesPoint); i++ {
		if i > position && (*candidatesPoint)[i] == (*candidatesPoint)[i - 1] {
			continue
		}
		search(resultsPoint, candidatesPoint, i + 1, append(current, (*candidatesPoint)[i]), target - (*candidatesPoint)[i])
	}
}
