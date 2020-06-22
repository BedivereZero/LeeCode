package algorithms

func letterCombinations(digits string) []string {
	var out []string
	for _, i := range digits {
		var cur, cs []string
		switch i {
		case '1':
			return cur
		case '2':
			cs = []string {"a", "b", "c"}
		case '3':
			cs = []string {"d", "e", "f"}
		case '4':
			cs = []string {"g", "h", "i"}
		case '5':
			cs = []string {"j", "k", "l"}
		case '6':
			cs = []string {"m", "n", "o"}
		case '7':
			cs = []string {"p", "q", "r", "s"}
		case '8':
			cs = []string {"t", "u", "v"}
		case '9':
			cs = []string {"w", "x", "y", "z"}
		}
		for _, c := range cs {
			if len(out) == 0 {
				cur = append(cur, c)
			}
			for _, e := range out {
				cur = append(cur, e + c)
			}
		}
		out = cur
	}
	return out
}
