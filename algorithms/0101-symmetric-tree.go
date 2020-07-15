package algorithms

func check(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == r
	}
	return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}
