package algorithms

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	p := root.Left
	if root.Left == nil {
		p = root
	} else {
		for p.Right != nil {
			p = p.Right
		}
	}
	root.Left, root.Right, p.Right = nil, root.Left, root.Right
}
