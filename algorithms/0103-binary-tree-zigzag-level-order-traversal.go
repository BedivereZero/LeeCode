package algorithms

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	if level%2 == 0 {
		(*traversal)[level] = append((*traversal)[level], root.Val)
	} else {
		(*traversal)[level] = append([]int{root.Val}, (*traversal)[level]...)
	}

	travel(root.Left, traversal, level+1)
	travel(root.Right, traversal, level+1)
}
