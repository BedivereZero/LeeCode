package algorithms

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func max(n ...int) int {
	for i := 1; i < len(n); i++ {
		if n[0] < n[i] {
			n[0] = n[i]
		}
	}
	return n[0]
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	var l, r int
	if root.Left != nil {
		l = root.Left.Val
	}
	if root.Right != nil {
		r = root.Right.Val
	}

	if abs(l-r) > 1 {
		return false
	}

	root.Val = max(l, r) + 1
	return true
}
