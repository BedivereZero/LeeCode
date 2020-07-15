package algorithms

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	prev, slow, fast := head, head.Next, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		prev = prev.Next
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	root := &TreeNode{Val: slow.Val}

	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)

	return root
}
