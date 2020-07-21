package algorithms

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connect(root.Left)
	connect(root.Right)

	for l, r := root.Left, root.Right; l != nil && r != nil; {
		p := l
		for p.Next != nil {
			p = p.Next
		}

		for l != nil {
			if l.Left != nil {
				l = l.Left
				break
			}
			if l.Right != nil {
				l = l.Right
				break
			}
			l = l.Next
		}

		p.Next = r

		for r != nil {
			if r.Left != nil {
				r = r.Left
				break
			}
			if r.Right != nil {
				r = r.Right
				break
			}
			r = r.Next
		}
	}
	return root
}
