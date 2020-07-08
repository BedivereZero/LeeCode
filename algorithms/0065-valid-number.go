package algorithms

func isNumber(s string) bool {
	start, end := 0, len(s)-1
	for ; start < len(s) && s[start] == ' '; start++ {
	}
	for ; end >= 0 && s[end] == ' '; end-- {
	}
	if start > end {
		return false
	}
	s = s[start : end+1]
	bPos, ePos := 0, 0
	for ; ePos < len(s) && s[ePos] != 'e'; ePos++ {
	}
	if s[bPos] == '+' || s[bPos] == '-' {
		bPos += 1
	}
	hasBase, hasExp := false, false
	hasPoint := false
	for ; bPos < ePos; bPos++ {
		if s[bPos] == '.' && !hasPoint {
			hasPoint = true
		} else if s[bPos] == '.' && hasPoint {
			return false
		} else if !(s[bPos] >= '0' && s[bPos] <= '9') {
			return false
		} else {
			hasBase = true
		}
	}
	if !hasBase {
		return false
	}
	pos := ePos + 1
	if pos < len(s) && (s[pos] == '+' || s[pos] == '-') {
		pos += 1
	}
	for ; pos < len(s); pos++ {
		if !(s[pos] >= '0' && s[pos] <= '9') {
			return false
		} else {
			hasExp = true
		}
	}
	if ePos < len(s) && !hasExp {
		return false
	}
	return true
}
