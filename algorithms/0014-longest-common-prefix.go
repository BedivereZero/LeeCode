func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	length := len(strs[0])
	for i := 1; i < len(strs); i++ {
		var compare_len int
		if length < len(strs[i]) {
			compare_len = length
		} else {
			compare_len = len(strs[i])
		}
		found := false
		for j := 0; j < compare_len; j++ {
			if strs[0][j] != strs[i][j] {
				length = j
				found = true
				break
			}
		}
		if !found {
			length = compare_len
		}
	}
	return strs[0][:length]
}
