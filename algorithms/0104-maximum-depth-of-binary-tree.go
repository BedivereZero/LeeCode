package algorithms

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	a, b := maxDepth(root.Left), maxDepth(root.Right)
	if a < b {
		return b + 1
	}
	return a + 1
}
