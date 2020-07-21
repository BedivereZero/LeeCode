package algorithms

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	for l, r := root.Left, root.Right; l != nil && r != nil; {
		l.Next = r
		l = l.Right
		r = r.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
