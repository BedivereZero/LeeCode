package algorithms

func reverseKGroup(head *ListNode, k int) *ListNode {
	root := &ListNode{
		Next: head,
	}
	cur := root
	for cur.Next != nil {
		end := cur
		len := 0
		for end.Next != nil && len < k {
			end = end.Next
			len++
		}
		if len < k {
			break
		}
		cur.Next = reverse(cur.Next, end.Next)
		for i := 0; i < k; i++ {
			cur = cur.Next
		}
	}
	return root.Next
}

func reverse(head, tail *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	new, cur, next := head, head, head.Next
	cur.Next = tail
	for next != tail {
		cur, next, next.Next = next, next.Next, cur
		new = cur
	}
	return new
}
