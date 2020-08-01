package algorithms

func reverseWords(s string) string {
	stringSlice := []string{}
	head, tail := 0, 0
	for ; head < len(s); head = tail {
		// find word begin
		for ; head < len(s) && s[head] == ' '; head++ {
		}
		// find word end
		for tail = head; tail < len(s) && s[tail] != ' '; tail++ {
		}
		if head < tail {
			stringSlice = append(stringSlice, s[head:tail])
		}
	}
	t := ""
	for i := len(stringSlice) - 1; i >= 0; i-- {
		if i < len(stringSlice)-1 {
			t += " "
		}
		t += stringSlice[i]
	}
	return t
}
