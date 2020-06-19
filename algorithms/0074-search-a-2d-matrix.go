package algorithms

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	x, found := binarySearchMatrix(matrix, target)
	if found || x < 1 {
		return found
	} else {
		x--
	}

	_, found = binarySearch(matrix[x], target)
	return found
}

func binarySearchMatrix(matrix [][]int, target int) (int, bool) {
	head, tail := 0, len(matrix)
	for head < tail {
		mid := (head + tail) / 2
		if matrix[mid][0] == target {
			return mid, true
		} else if matrix[mid][0] < target {
			head = mid + 1
		} else {
			tail = mid
		}
	}
	return head, head < len(matrix) && matrix[head][0] == target
}

func binarySearch(array []int, target int) (int, bool) {
	head, tail := 0, len(array)
	for head < tail {
		mid := (head + tail) / 2
		if array[mid] == target {
			return mid, true
		} else if array[mid] < target {
			head = mid + 1
		} else {
			tail = mid
		}
	}
	return head, head < len(array) && array[head] == target
}
