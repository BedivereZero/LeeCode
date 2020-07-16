package algorithms

func max(n ...int) int {
	for i := 1; i < len(n); i++ {
		if n[0] < n[i] {
			n[0] = n[i]
		}
	}
	return n[0]
}

func maxPathSum(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	if root.Left == nil {
		l := maxPathSum(root.Right)
		if root.Right.Val > 0 {
			root.Val += root.Right.Val
		}
		return max(root.Val, l)
	}
	if root.Right == nil {
		r := maxPathSum(root.Left)
		if root.Left.Val > 0 {
			root.Val += root.Left.Val
		}
		return max(root.Val, r)
	}

	l := maxPathSum(root.Left)
	r := maxPathSum(root.Right)
	v := root.Val
	root.Val += max(root.Left.Val, root.Right.Val, 0)
	return max(root.Left.Val+v, root.Left.Val+v+root.Right.Val, v+root.Right.Val, v, l, r)
}
