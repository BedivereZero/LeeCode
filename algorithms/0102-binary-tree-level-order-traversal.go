package algorithms

func levelOrder(root *TreeNode) [][]int {
	traversal := [][]int{}
	travel(root, &traversal, 0)
	return traversal
}

func travel(root *TreeNode, traversal *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*traversal) < level+1 {
		*traversal = append(*traversal, []int{})
	}
	(*traversal)[level] = append((*traversal)[level], root.Val)

	travel(root.Left, traversal, level+1)
	travel(root.Right, traversal, level+1)
}
