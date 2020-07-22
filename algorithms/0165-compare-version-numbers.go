package algorithms

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1s := strings.Split(version1, ".")
	v2s := strings.Split(version2, ".")
	for i := 0; i < len(v1s) && i < len(v2s); i++ {
		v1, _ := strconv.Atoi(v1s[i])
		v2, _ := strconv.Atoi(v2s[i])
		if v1 < v2 {
			return -1
		}
		if v1 > v2 {
			return 1
		}
	}
	for i := len(v1s); i < len(v2s); i++ {
		if v, _ := strconv.Atoi(v2s[i]); v > 0 {
			return -1
		}
	}
	for i := len(v2s); i < len(v1s); i++ {
		if v, _ := strconv.Atoi(v1s[i]); v > 0 {
			return 1
		}
	}
	return 0
}
