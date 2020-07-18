package algorithms

func convertToTitle(n int) string {
	s := ""
	for true {
		n--
		s = string([]byte{'A' + byte(n%26)}) + s
		n = n / 26
		if n == 0 {
			break
		}
	}
	return s
}
