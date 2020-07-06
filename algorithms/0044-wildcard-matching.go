package algorithms

func isMatch(s string, p string) bool {
	// create DP matrix
	f := make([][]bool, len(p)+1)
	for x := range f {
		f[x] = make([]bool, len(s)+1)
	}

	// if s or p is empty
	f[0][0] = true
	for x := 1; x <= len(p); x++ {
		f[x][0] = p[x-1] == '*' && f[x-1][0]
	}

	// calculate DP matrix
	for x := 1; x <= len(p); x++ {
		for y := 1; y <= len(s); y++ {
			switch p[x-1] {
			case s[y-1], '?':
				f[x][y] = f[x-1][y-1]
			case '*':
				for i := 0; !f[x][y] && i <= y; i++ {
					f[x][y] = f[x-1][i]
				}
			}
		}
	}
	return f[len(p)][len(s)]
}
