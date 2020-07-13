package algorithms

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) == 0 {
		return len(s3) == 0
	}
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}

	if s3[0] == s1[0] && isInterleave(s1[1:], s2, s3[1:]) {
		return true
	}
	if s3[0] == s2[0] && isInterleave(s1, s2[1:], s3[1:]) {
		return true
	}
	return false
}
