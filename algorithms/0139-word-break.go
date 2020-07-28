package algorithms

func wordBreak(s string, wordDict []string) bool {
	record := make([]bool, len(s))
	for i := range s {
		for _, word := range wordDict {
			if record[i] {
				break
			}
			if i+1 < len(word) {
				continue
			}
			record[i] = s[i+1-len(word):i+1] == word
			if i+1 > len(word) {
				record[i] = record[i] && record[i-len(word)]
			}
		}
	}
	return record[len(record)-1]
}
