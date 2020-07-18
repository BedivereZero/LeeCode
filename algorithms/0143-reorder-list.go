package algorithms

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	stack := &Stack{}
	for cur := slow.Next; cur != nil; {
		stack.push(cur)
		cur = cur.Next
	}

	slow.Next = nil

	for cur := head; stack.depth > 0; {
		n := stack.pop().(*ListNode)
		cur.Next, n.Next = n, cur.Next
		cur = cur.Next.Next
	}
}
