package algorithms

func findRepeatedDnaSequences(s string) []string {
	L := 10
	if len(s) < L {
		return []string{}
	}

	record := map[uint32]int{}
	result := []string{}

	hash := hash10(s)
	record[hash] = 1
	for i := 0; i+L < len(s); i++ {
		hash = rotateHash(hash, s[i+L])
		switch record[hash] {
		case 0:
			record[hash] = 1
		case 1:
			result = append(result, s[i+1:i+L+1])
			record[hash] = 2
		}
	}
	return result
}

func hash10(s string) uint32 {
	var h uint32
	for i := 0; i < len(s) && i < 10; i++ {
		h = h << 2
		switch s[i] {
		case 'A':
			h = h ^ 0
		case 'C':
			h = h ^ 1
		case 'G':
			h = h ^ 2
		case 'T':
			h = h ^ 3
		}
	}
	return h
}

func rotateHash(h uint32, b byte) uint32 {
	h = (h << 2) % (1 << 20)
	switch b {
	case 'A':
		h = h ^ 0
	case 'C':
		h = h ^ 1
	case 'G':
		h = h ^ 2
	case 'T':
		h = h ^ 3
	}
	return h
}
