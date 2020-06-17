package algorithms

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	var result [][]int
	quickSort(intervals, 0, len(intervals))
	current := intervals[0]
	for _, v := range intervals[1:] {
		if current[1] < v[0] {
			result = append(result, current)
			current = v
		} else if current[1] < v[1] {
			current[1] = v[1]
		}
	}
	return append(result, current)
}

func quickSort(intervals [][]int, head int, tail int) {
	if head+1 < tail {
		pivot := partition(intervals, head, tail)
		quickSort(intervals, head, pivot)
		quickSort(intervals, pivot+1, tail)
	}
}

func partition(intervals [][]int, head int, tail int) int {
	pivot := intervals[tail-1]
	i := head
	for j := head; j+1 < tail; j++ {
		if compare(intervals[j], pivot) {
			intervals[i], intervals[j] = intervals[j], intervals[i]
			i++
		}
	}
	intervals[i], intervals[tail-1] = intervals[tail-1], intervals[i]
	return i
}

func compare(interval0 []int, interval1 []int) bool {
	if interval0[0] == interval1[0] {
		return interval0[1] < interval1[1]
	} else {
		return interval0[0] < interval1[0]
	}
}
