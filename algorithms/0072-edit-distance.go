package algorithms

func minDistance(word1 string, word2 string) int {
	if len(word1) < len(word2) {
		word1, word2 = word2, word1
	}
	f := make([]int, len(word2)+1)

	for x := 0; x < len(word1)+1; x++ {
		var pre int
		for y := range f {
			if x == 0 {
				f[y] = y
				continue
			}
			if y == 0 {
				pre, f[0] = f[0], x
				continue
			}
			if word1[x-1] != word2[y-1] {
				pre++
			}
			if pre > f[y]+1 {
				pre = f[y] + 1
			}
			if pre > f[y-1]+1 {
				pre = f[y-1] + 1
			}
			pre, f[y] = f[y], pre
		}
	}
	return f[len(word2)]
}
