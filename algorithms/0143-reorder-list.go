package algorithms

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	prev, tail := head, head.Next
	for tail.Next != nil {
		prev = prev.Next
		tail = tail.Next
	}

	head.Next, prev.Next, tail.Next = tail, nil, head.Next
	reorderList(head.Next.Next)
}
