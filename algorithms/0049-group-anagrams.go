package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string {"eat", "tea", "tan", "ate", "nat", "bat"}
	// s := []string {"cab","tin","pew","duh","may","ill","buy","bar","max","doc"}
	for _, v := range groupAnagrams(s) {
		fmt.Println(v)
	}
}

func groupAnagrams(strs []string) [][]string {
	nums := make([]int, len(strs))
	for i, v := range strs {
		nums[i] = charToDecimal(v)
	}
	quicksort(strs, nums, 0, len(strs))
	var result [][]string
	var current []string
	for i, v := range strs {
		if i > 0 && nums[i] != nums[i - 1] {
			result = append(result, current)
			current = []string {}
		}
		current = append(current, v)
	}
	if len(current) > 0 {
		result = append(result, current)
	}
	return result
}

func charToDecimal(chars string) int {
	temp := make([]int, len(chars))
	for i, v := range chars {
		temp[i] = int(v - 'a')
	}
	sort.Ints(temp)
	sum, base := 0, 1
	for _, v := range temp {
		sum += int(v - 'a') * base
		base *= 26
	}
	return sum
}

func quicksort(strs []string, nums []int, head int, tail int) {
	if head + 1 < tail {
		pivot := partition(strs, nums, head, tail)
		quicksort(strs, nums, head, pivot)
		quicksort(strs, nums, pivot + 1, tail)
	}
}

func partition(strs []string, nums []int, head int, tail int) int {
	pivot := nums[tail - 1]
	i := head
	for j := head; j + 1 < tail; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			strs[i], strs[j] = strs[j], strs[i]
			i++
		}
	}
	nums[i], nums[tail - 1] = nums[tail - 1], nums[i]
	strs[i], strs[tail - 1] = strs[tail - 1], strs[i]
	return i
}
