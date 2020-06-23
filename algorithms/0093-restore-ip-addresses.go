package algorithms

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return []string{}
	}
	addresses := []string{}
	current := []uint8{}
	if s[0] == '0' {
		search(&addresses, append(current, uint8(0)), s[1:])
		return addresses
	}
	for i := 1; i <= len(s); i++ {
		if num, err := strconv.ParseUint(s[:i], 10, 8); err == nil {
			search(&addresses, append(current, uint8(num)), s[i:])
		}
	}
	return addresses
}

func search(addresses *[]string, current []uint8, s string) {
	if len(current) == 4 && len(s) == 0 {
		*addresses = append(*addresses, stringFromIntArray(current))
		return
	}
	if len(current) == 4 || len(s) == 0 {
		return
	}
	if s[0] == '0' {
		search(addresses, append(current, uint8(0)), s[1:])
		return
	}
	for i := 1; i <= len(s); i++ {
		if num, err := strconv.ParseUint(s[:i], 10, 8); err == nil {
			search(addresses, append(current, uint8(num)), s[i:])
		}
	}
}

func stringFromIntArray(array []uint8) string {
	var s string
	for i := 0; i < len(array); i++ {
		if i > 0 {
			s = s + "."
		}
		s += strconv.FormatUint(uint64(array[i]), 10)
	}
	return s
}
