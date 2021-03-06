package algorithms

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	i := 0
	for ; i < len(inorder); i++ {
		if root.Val == inorder[i] {
			break
		}
	}

	root.Left = buildTree(preorder[1:1+i], inorder[:i])
	root.Right = buildTree(preorder[1+i:], inorder[i+1:])

	return root
}
