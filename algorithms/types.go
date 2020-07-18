package algorithms

// Node is a node
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ListNode is a list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type (
	node struct {
		val  interface{}
		next *node
	}
	Stack struct {
		top   *node
		depth int
	}
)

func (s *Stack) push(element interface{}) {
	n := &node{val: element}
	if s.depth == 0 {
		s.top = n
	} else {
		n.next = s.top
		s.top = n
	}
	s.depth++
}

func (s *Stack) pop() interface{} {
	if s.depth == 0 {
		return nil
	}
	n := s.top
	if s.depth == 1 {
		s.top = nil
	} else {
		s.top = s.top.next
	}
	s.depth--
	return n.val
}
