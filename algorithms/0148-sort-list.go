package algorithms

func sortList(head *ListNode) *ListNode {
	head, _ = quickSortList(head)
	return head
}

func quickSortList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	pivot := head

	aHalf, bHalf := splitList(head.Next, func(n *ListNode) bool {
		return n.Val < pivot.Val
	})
	aHead, aTail := quickSortList(aHalf)
	bHead, bTail := quickSortList(bHalf)

	if aHalf == nil {
		head = pivot
	} else {
		head = aHead
		aTail.Next = pivot
	}

	var tail *ListNode
	if bHalf == nil {
		tail = pivot
		pivot.Next = nil
	} else {
		tail = bTail
		pivot.Next = bHead
	}

	return head, tail
}

func splitList(head *ListNode, f func(*ListNode) bool) (*ListNode, *ListNode) {
	aRoot, bRoot := ListNode{}, ListNode{}
	aCurr, bCurr := &aRoot, &bRoot
	for pos := head; pos != nil; pos = pos.Next {
		if f(pos) {
			aCurr.Next = pos
			aCurr = aCurr.Next
		} else {
			bCurr.Next = pos
			bCurr = bCurr.Next
		}
	}
	aCurr.Next, bCurr.Next = nil, nil
	return aRoot.Next, bRoot.Next
}
