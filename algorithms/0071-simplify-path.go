package algorithms

func split(path string) []string {
	head, tail, hierarchies := 1, 1, []string {}
	for tail < len(path) {
		if path[tail] == '/' {
			if head < tail {
				hierarchies = append(hierarchies, path[head:tail])
			}
			head, tail = tail + 1, tail + 1
		} else {
			tail++
		}
	}
	return hierarchies
}

func simplify(hierarchies []string) string {
	index := 0
	for _, value := range hierarchies {
		switch value {
		case ".":
			continue
		case "..":
			if index > 0 {
				index--
			}
		default:
			hierarchies[index] = value
			index ++
		}
	}
	if index == 0 {
		return "/"
	}
	result := ""
	for _, value := range hierarchies[:index] {
		result += "/" + value
	}
	return result
}
