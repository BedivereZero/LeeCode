package algorithms

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	t := []int{root.Val}
	t = append(t, preorderTraversal(root.Left)...)
	t = append(t, preorderTraversal(root.Right)...)
	return t
}
