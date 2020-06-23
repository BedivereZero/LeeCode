package algorithms

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Element struct {
	Node *TreeNode
	Step int
}

type Stack struct {
	top *node
	length int
}

type node struct {
	value *Element
	prev *node
}

func newStack() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Len() int {
	return this.length
}

func (this *Stack) Peek() *Element {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

func (this *Stack) Pop() *Element {
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value *Element) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

func inorderTraversal(root *TreeNode) []int {
	var footprint []int
	stack := newStack()
	cur := &Element{
		Node: root,
		Step: 0,
	}
	for stack.Len() != 0 || cur != nil {
		if cur.Node == nil {
			fmt.Printf("A: cur:nil, step:%2d\n", cur.Step)
		} else {
			fmt.Printf("A: cur:%3d, step:%2d\n", cur.Node.Val, cur.Step)
		}
		switch cur.Step {
		case 0:
			if cur.Node == nil {
				// current is nil
				cur = stack.Pop()
				if cur != nil {
					cur.Step++
				}
			} else {
				// traversal left
				stack.Push(cur)
				cur = &Element{
					Node: cur.Node.Left,
					Step: 0,
				}
			}
		case 1:
			// record
			footprint = append(footprint, cur.Node.Val)
			// traversal right
			stack.Push(cur)
			cur = &Element{
				Node: cur.Node.Right,
				Step: 0,
			}
		case 2:
			cur = stack.Pop()
			if cur != nil {
				cur.Step++
			}
		}
		fmt.Printf("B: %p\n", cur)
	}
	return footprint
}

func traversalRecursively(footprint *[]int, root *TreeNode) {
	// 0 check node is nil and traversal left
	if root == nil {
		return
	}
	traversalRecursively(footprint, root.Left)
	// 1 traversal right
	*footprint = append(*footprint, root.Val)
	traversalRecursively(footprint, root.Right)
	// 2 finish travel
}
