package algorithms

func rightSideView(root *TreeNode) []int {
	view := []int{}
	currentLevel := []*TreeNode{root}
	for len(currentLevel) > 0 {
		if n := getLastNotNull(currentLevel); n != nil {
			view = append(view, n.Val)
		}
		nextLevel := []*TreeNode{}
		for _, node := range currentLevel {
			if node == nil {
				continue
			}
			nextLevel = append(nextLevel, node.Left)
			nextLevel = append(nextLevel, node.Right)
		}
		currentLevel = nextLevel
	}
	return view
}

func getLastNotNull(nodes []*TreeNode) *TreeNode {
	for i := len(nodes) - 1; i >= 0; i-- {
		if nodes[i] != nil {
			return nodes[i]
		}
	}
	return nil
}
