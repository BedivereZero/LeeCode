package algorithms

type (
	TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	node struct {
		treeNode *TreeNode
		phase    int
		Next     *node
	}
	stack struct {
		top   *node
		depth int
	}
)

func (s *stack) Push(tn *TreeNode, p int) {
	n := &node{treeNode: tn, phase: p, Next: s.top}
	s.top = n
	s.depth++
}

func (s *stack) Pop() (*TreeNode, int) {
	n := s.top
	if n == nil {
		return nil, 0
	}
	s.top = s.top.Next
	s.depth--
	return n.treeNode, n.phase
}

func (s *stack) Depth() int {
	return s.depth
}

func postorderTraversal(root *TreeNode) []int {
	r := []int{}
	record := &stack{}
	curr, phase := root, 0
	for curr != nil || record.Depth() > 0 {
		switch phase {
		case 0:
			if curr == nil {
				curr, phase = record.Pop()
				continue
			}
			record.Push(curr, phase+1)
			curr, phase = curr.Left, 0
		case 1:
			record.Push(curr, phase+1)
			curr, phase = curr.Right, 0
		case 2:
			r = append(r, curr.Val)
			curr, phase = record.Pop()
		}
	}
	return r
}
