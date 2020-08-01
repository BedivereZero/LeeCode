package algorithms

import "strings"

func reverseWords(s string) string {
	fileds := strings.Fields(s)
	for i, j := 0, len(fileds)-1; i < j; i, j = i+1, j-1 {
		fileds[i], fileds[j] = fileds[j], fileds[i]
	}
	return strings.Join(fileds, " ")
}
