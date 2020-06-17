// package algorithms

package main

import (
	"fmt"
	"bytes"
)

func main() {
	n, k := 4, 9
	fmt.Println(getPermutation(n, k))
}

func getPermutation(n int, k int) string {
	permutationAsList := make([]int, n)

	// Generate factorials
	factorials := getFactorials(n)

	// Generate candidates
	candidates := getCandidates(n)

	// Transform index 1...k to 0...K
	K := k - 1
	for i := 0; i < n; i++ {
		position := K / factorials[len(candidates) - 1]
		permutationAsList[i] = candidates[position]
		K = K % factorials[len(candidates) - 1]
		candidates = append(candidates[:position], candidates[position + 1:]...)
	}

	// Transform []int to string
	var buffer bytes.Buffer
	for _, v := range permutationAsList {
		buffer.WriteByte(byte(v + '0'))
	}
	return buffer.String()
}

func getFactorials(n int) []int {
	factorials := make([]int, n)
	factorials[0] = 1
	for i := 1; i < n; i ++ {
		factorials[i] = factorials[i - 1] * i
	}
	return factorials
}

func getCandidates(n int) []int {
	candidates := make([]int, n)
	for i := 0; i < n; i++ {
		candidates[i] = i + 1
	}
	return candidates
}
