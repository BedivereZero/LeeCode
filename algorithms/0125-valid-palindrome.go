package algorithms

import "strings"

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for ; l < len(s) && (s[l] < 'a' || s[l] > 'z') && (s[l] < '0' || s[l] > '9'); l++ {
	}
	for ; r >= 0 && (s[r] < 'a' || s[r] > 'z') && (s[r] < '0' || s[r] > '9'); r-- {
	}
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
		for ; l < len(s) && (s[l] < 'a' || s[l] > 'z') && (s[l] < '0' || s[l] > '9'); l++ {
		}
		for ; r >= 0 && (s[r] < 'a' || s[r] > 'z') && (s[r] < '0' || s[r] > '9'); r-- {
		}
	}
	return true
}
