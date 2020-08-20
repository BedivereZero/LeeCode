package algorithms

type slope struct {
	dx int
	dy int
	ho bool
}

func calcGCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for a*b > 0 {
		if a < b {
			a, b = b, a
		}
		a, b = b, a%b
	}
	return a
}

func maxPoints(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	max := 0
	// record all lines over two nodes
	for i := 0; i < len(points); i++ {
		line := make(map[slope]int)
		same := 0
		maxPart := 0
		for j := i + 1; j < len(points); j++ {
			// record
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]
			if dx == 0 && dy == 0 {
				same++
				continue
			}
			s := slope{}
			if dx != 0 && dy != 0 {
				gcd := calcGCD(dx, dy)
				s.dx = dx / gcd
				s.dy = dy / gcd
				if s.dx < 0 {
					s.dx = -s.dx
					s.dy = -s.dy
				}
			} else {
				s.ho = dy == 0
			}
			line[s]++
			if maxPart < line[s] {
				maxPart = line[s]
			}
		}
		if max < maxPart+same {
			max = maxPart + same
		}
	}
	return max + 1
}
