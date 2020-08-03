package algorithms

type (
	node struct {
		Val  interface{}
		Next *node
	}
	Stack struct {
		top   *node
		depth int
	}
	BSTIterator struct {
		stack *Stack
	}
)

func (s *Stack) Push(value interface{}) {
	n := &node{Val: value}
	n.Next = s.top
	s.top = n
	s.depth++
}

func (s *Stack) Pop() interface{} {
	if s.depth == 0 {
		return nil
	}

	n := s.top
	s.top = s.top.Next
	s.depth--
	return n.Val
}

func (s *Stack) Depth() int {
	return s.depth
}

func Constructor(root *TreeNode) BSTIterator {
	i := BSTIterator{
		stack: &Stack{},
	}
	i.pushLeftSubTree(root)
	return i
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	n := this.stack.Pop().(*TreeNode)
	this.pushLeftSubTree(n.Right)
	return n.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.stack.Depth() > 0
}

func (this *BSTIterator) pushLeftSubTree(root *TreeNode) {
	for cur := root; cur != nil; cur = cur.Left {
		this.stack.Push(cur)
	}
}
