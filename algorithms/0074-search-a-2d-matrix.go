package algorithms

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	index := binarySearchMatrix(matrix, target)
	if index/len(matrix[0]) == len(matrix) {
		return false
	} else {
		return target == matrix[index/len(matrix[0])][index%len(matrix[0])]
	}
}

func binarySearchMatrix(matrix [][]int, target int) int {
	head, tail := 0, len(matrix)*len(matrix[0])
	for head < tail {
		mid := (head + tail) / 2
		if matrix[mid/len(matrix[0])][mid%len(matrix[0])] == target {
			return mid
		} else if matrix[mid/len(matrix[0])][mid%len(matrix[0])] < target {
			head = mid + 1
		} else {
			tail = mid
		}
	}
	return head
}
