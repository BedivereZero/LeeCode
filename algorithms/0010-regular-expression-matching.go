package algorithms

func isMatch(s string, p string) bool {
	// Create 2D Matrix f
	f := make([][]bool, len(p)+1)
	for i := range f {
		f[i] = make([]bool, len(s)+1)
	}

	f[0][0] = true
	for x := range f {
		for y := range f[x] {
			if x == 0 {
				continue
			}
			switch p[x-1] {
			case '.':
				if y == 0 {
					continue
				}
				f[x][y] = f[x-1][y-1]
			case '*':
				if x < 2 || p[x-2] == '*' {
					continue
				}
				for i := 0; !f[x][y] && i < y+1; i++ {
					if p[x-2] == '.' {
						f[x][y] = f[x-2][y-i]
					} else {
						f[x][y] = f[x-2][y-i] && single(s[y-i:y], p[x-2])
					}
				}
			default:
				if y == 0 {
					continue
				}
				f[x][y] = f[x-1][y-1] && p[x-1] == s[y-1]
			}
		}
	}

	return f[len(p)][len(s)]
}

func single(s string, c byte) bool {
	for i := range s {
		if s[i] != c {
			return false
		}
	}
	return true
}
