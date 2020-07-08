package algorithms

func insert(intervals [][]int, newInterval []int) [][]int {
	// find left, left[1] == newInterval[0] OR left[l] is Minimum value greater than newInterval[0]
	l := search(len(intervals), func(i int) bool { return intervals[i][1] >= newInterval[0] })

	// find right, left[0] = newInterval[1] OR right[0] is Maximum value less than newInteravl[1]
	r := search(len(intervals), func(i int) bool { return intervals[i][0] > newInterval[1] })

	// calculate newInterval
	if l < r {
		if l < len(intervals) && intervals[l][0] < newInterval[0] {
			newInterval[0] = intervals[l][0]
		}
		if r > 0 && intervals[r-1][1] > newInterval[1] {
			newInterval[1] = intervals[r-1][1]
		}
	}

	// create new
	newIntervals := make([][]int, 1+len(intervals)+l-r)
	copy(newIntervals, intervals[:l])
	newIntervals[l] = make([]int, 2)
	copy(newIntervals[l], newInterval)
	copy(newIntervals[l+1:], intervals[r:])
	return newIntervals
}

func search(n int, f func(int) bool) int {
	h, t := 0, n
	for h < t {
		m := (h + t) / 2
		if !f(m) {
			h = m + 1
		} else {
			t = m
		}
	}
	return h
}
