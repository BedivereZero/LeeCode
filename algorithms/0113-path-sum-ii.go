package algorithms

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	return travel(root, sum, []int{})
}

func travel(root *TreeNode, sum int, path []int) [][]int {
	r := [][]int{}
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			n := make([]int, len(path))
			copy(n, path)
			n = append(n, root.Val)
			r = append(r, n)
		}
		return r
	}

	if root.Left != nil {
		r = append(r, travel(root.Left, sum-root.Val, append(path, root.Val))...)
	}
	if root.Right != nil {
		r = append(r, travel(root.Right, sum-root.Val, append(path, root.Val))...)
	}
	return r
}
