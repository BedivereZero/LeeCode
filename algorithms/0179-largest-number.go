package algorithms

import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	fields := make([]string, len(nums))
	for i, n := range nums {
		fields[i] = fmt.Sprintf("%d", n)
	}
	sort.Slice(fields, func(i, j int) bool {
		a, b := fields[i]+fields[j], fields[j]+fields[i]
		for p := 0; p < len(a); p++ {
			if a[p] != b[p] {
				return int(a[p]-'0') < int(b[p]-'0')
			}
		}
		return true
	})
	for i, j := 0, len(fields)-1; i < j; i, j = i+1, j-1 {
		fields[i], fields[j] = fields[j], fields[i]
	}
	s := strings.Join(fields, "")
	if s[0] == '0' {
		return "0"
	}
	return s
}
