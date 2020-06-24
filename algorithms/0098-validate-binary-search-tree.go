package algorithms

import "errors"

func isValidBST(root *TreeNode) bool {
	_, _, err := edge(root)
	if err == nil {
		return true
	} else {
		return false
	}

}

// get valid bst and return min and max node
func edge(root *TreeNode) (*TreeNode, *TreeNode, error) {
	if root == nil {
		return nil, nil, nil
	}

	var l, r *TreeNode

	// valid left
	ll, lr, err := edge(root.Left)
	if err != nil {
		return nil, nil, err
	}
	if lr == nil {
		l = root
	} else if lr.Val < root.Val {
		l = ll
	} else {
		return nil, nil, errors.New("invalid")
	}

	// valid right
	rl, rr, err := edge(root.Right)
	if err != nil {
		return nil, nil, err
	}
	if rl == nil {
		r = root
	} else if root.Val < rl.Val {
		r = rr
	} else {
		return nil, nil, errors.New("invalid")
	}

	return l, r, nil
}
