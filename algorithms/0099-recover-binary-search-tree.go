package algorithms

// orientation, true: head -> tail, false, tail -> head
func travel(orientation bool, root *TreeNode, cache []*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var a, b *TreeNode
	if orientation {
		a, b = root.Left, root.Right
	} else {
		a, b = root.Right, root.Left
	}

	if tn := travel(orientation, a, cache); tn != nil {
		return tn
	}

	if cache[0] != nil && (orientation && cache[0].Val > root.Val || !orientation && cache[0].Val < root.Val) {
		return cache[0]
	}
	cache[0] = root

	if tn := travel(orientation, b, cache); tn != nil {
		return tn
	}

	return nil
}

func recoverTree(root *TreeNode) {
	cache := []*TreeNode{nil}

	// travel from head
	a := travel(true, root, cache)

	// clean cache
	cache[0] = nil

	// travel from tail
	b := travel(false, root, cache)

	a.Val, b.Val = b.Val, a.Val
}
