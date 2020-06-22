package algorithms

func romanToInt(s string) int {
	cha := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s)-1; i++ {
		if cha[s[i]] < cha[s[i+1]] {
			result -= cha[s[i]]
		} else {
			result += cha[s[i]]
		}
	}
	result += cha[s[len(s)-1]]
	return result
}
