package algorithms

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return root.Left != nil && hasPathSum(root.Left, sum-root.Val) || root.Right != nil && hasPathSum(root.Right, sum-root.Val)
}
