package algorithms

func min(n ...int) int {
	for i := 1; i < len(n); i++ {
		if n[i] < n[0] {
			n[0] = n[i]
		}
	}
	return n[0]
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
