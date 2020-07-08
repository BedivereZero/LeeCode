package algorithms

func minDistance(word1 string, word2 string) int {
	f := make([][]int, len(word1)+1)
	for x := range f {
		f[x] = make([]int, len(word2)+1)
	}
	for x := range f {
		for y := range f[x] {
			if x == 0 || y == 0 {
				f[x][y] = x + y
				continue
			}
			f[x][y] = f[x-1][y-1]
			if word1[x-1] != word2[y-1] {
				f[x][y]++
			}
			if f[x][y] > f[x-1][y]+1 {
				f[x][y] = f[x-1][y] + 1
			}
			if f[x][y] > f[x][y-1]+1 {
				f[x][y] = f[x][y-1] + 1
			}
		}
	}
	return f[len(word1)][len(word2)]
}
