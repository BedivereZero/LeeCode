package algorithms

func strStr(haystack string, needle string) int {
	lh, ln := len(haystack), len(needle)
	for i := 0; i < lh-ln+1; i++ {
		if haystack[i:i+ln] == needle {
			return i
		}
	}
	return -1
}
