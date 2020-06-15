package main

import "fmt"

func main() {
	candidates := []int {2, 3, 5}
	target := 8
	for _, each := range combinationSum(candidates, target) {
		fmt.Println(each)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	return search(candidates, target, []int {})
}

func search(candidates []int, target int, current []int) [][]int {
	if target == 0 {
		new := make([]int, len(current))
		for i, e := range current {
			new[i] = e
		}
		return [][]int {new}
	}
	if target < 0 {
		return [][]int {}
	}
	var result [][]int
	for point, each := range candidates {
		cur := search(candidates[point:], target - each, append(current, each))
		result = append(result, cur...)
	}
	return result
}
