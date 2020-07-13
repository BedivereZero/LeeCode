package algorithms

import "log"

func check(original, target map[byte]int) bool {
	for k, v := range original {
		if v != target[k] {
			return false
		}
	}
	return true
}

func isScramble(s1 string, s2 string) bool {
	log.Printf("Chekck: %s, %s", s1, s2)
	l := len(s1)
	switch l {
	case 0:
		return false
	case 1:
		return s1[0] == s2[0]
	case 2:
		return s1[0] == s2[0] && s1[1] == s2[1] || s1[0] == s2[1] && s1[1] == s2[0]
	}

	// record appearance times
	orig := make(map[byte]int)
	head := make(map[byte]int)
	tail := make(map[byte]int)

	for i := 1; i < len(s1); i++ {
		orig[s1[i-1]]++
		head[s2[i-1]]++
		tail[s2[l-i]]++
		log.Println(orig, head, tail)
		if check(orig, head) && isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if check(orig, tail) && isScramble(s1[:i], s2[l-i:]) && isScramble(s1[i:], s2[:l-i]) {
			return true
		}
	}
	return false
}
