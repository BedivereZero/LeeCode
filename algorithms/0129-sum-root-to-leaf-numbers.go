package algorithms

func sumNumbers(root *TreeNode) int {
	return sumTree(root, 0)
}

func sumTree(root *TreeNode, pre int) int {
	if root == nil {
		return 0
	}

	pre = pre*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return pre
	}
	return sumTree(root.Left, pre) + sumTree(root.Right, pre)
}
