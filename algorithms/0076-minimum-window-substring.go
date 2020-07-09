package algorithms

func minWindow(s string, t string) string {
	timeActual, timeExpect := make(map[byte]int), make(map[byte]int)
	for i := range t {
		timeActual[t[i]] = 0
		timeExpect[t[i]]++
	}

	substring, exists := s, false
	for head, tail := 0, 0; tail < len(s); {
		// expand
		for ; !checkAllExist(timeExpect, timeActual) && tail < len(s); tail++ {
			if _, ok := timeActual[s[tail]]; ok {
				timeActual[s[tail]]++
			}
		}
		// reduce
		for ; checkAllExist(timeExpect, timeActual) && head < tail; head++ {
			if _, ok := timeActual[s[head]]; ok {
				timeActual[s[head]]--
			}
		}
		if head > 0 && timeActual[s[head-1]] < timeExpect[s[head-1]] {
			exists = true
			if tail-head+1 < len(substring) {
				substring = s[head-1 : tail]
			}
		}
	}

	if exists {
		return substring
	} else {
		return ""
	}
}

func checkAllExist(expect, actual map[byte]int) bool {
	for k, v := range expect {
		if actual[k] < v {
			return false
		}
	}
	return true
}
