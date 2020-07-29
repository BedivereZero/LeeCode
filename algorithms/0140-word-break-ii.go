package algorithms

func wordBreak(s string, wordDict []string) []string {
	recordString := make([][]string, len(s)+1)
	recordBool := make([]bool, len(s)+1)
	return traversal(s, wordDict, recordString, recordBool, 0)
}

func traversal(s string, wordDict []string, recordString [][]string, recordBool []bool, start int) []string {
	if recordBool[start] {
		return recordString[start]
	}

	if start == len(s) {
		recordString[start] = append(recordString[start], "")
	}
	for end := start + 1; end <= len(s); end++ {
		for _, word := range wordDict {
			if s[start:end] != word {
				continue
			}
			for _, l := range traversal(s, wordDict, recordString, recordBool, end) {
				if l == "" {
					recordString[start] = append(recordString[start], word)
				} else {
					recordString[start] = append(recordString[start], word+" "+l)
				}
			}
			break
		}
	}
	recordBool[start] = true
	return recordString[start]
}
