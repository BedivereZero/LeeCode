package algorithms

func levelOrderBottom(root *TreeNode) [][]int {
	traversal := [][]int{}
	travel(root, &traversal, 0)
	return traversal
}

func travelBottom(root *TreeNode, traversal *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*traversal) < level+1 {
		*traversal = append([][]int{{}}, *traversal...)
	}
	(*traversal)[len(*traversal)-1-level] = append((*traversal)[len(*traversal)-1-level], root.Val)

	travel(root.Left, traversal, level+1)
	travel(root.Right, traversal, level+1)
}
