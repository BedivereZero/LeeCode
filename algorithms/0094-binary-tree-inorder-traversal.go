package algorithms

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var footprint []int
	traversal(&footprint, root)
	return footprint
}

func traversal(footprint *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	traversal(footprint, root.Left)
	*footprint = append(*footprint, root.Val)
	traversal(footprint, root.Right)
}
