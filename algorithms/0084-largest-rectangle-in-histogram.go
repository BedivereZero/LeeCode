package algorithms

type (
	stack struct {
		top   *node
		depth int
	}
	node struct {
		value interface{}
		next  *node
	}
)

func (s *stack) push(element interface{}) {
	n := &node{value: element, next: s.top}
	s.top = n
	s.depth++
}

func (s *stack) pop() interface{} {
	if s.depth == 0 {
		return nil
	}
	n := s.top
	s.top = s.top.next
	s.depth--
	return n.value
}

func (s *stack) peek() interface{} {
	if s.depth == 0 {
		return nil
	}
	return s.top.value
}

func largestRectangleArea(heights []int) int {
	index := make([]int, len(heights))
	cache := make([]int, len(heights))
	depth := 0

	largest := 0
	for i, h := range heights {
		for depth > 0 && heights[index[depth-1]] > h {
			depth--
			if s := heights[index[depth]] * (i - index[depth] - 1); largest < s+cache[depth] {
				largest = s + cache[depth]
			}
		}
		width := i + 1
		if depth > 0 {
			width -= index[depth-1] + 1
		}
		s := width * h
		if depth > 0 && heights[index[depth-1]] == h {
			s += cache[depth-1]
			depth--
		}
		index[depth] = i
		cache[depth] = s
		depth++
		if largest < s {
			largest = s
		}
	}
	for depth > 0 {
		depth--
		if s := heights[index[depth]] * (len(heights) - index[depth] - 1); largest < s+cache[depth] {
			largest = s + cache[depth]
		}
	}
	return largest
}
