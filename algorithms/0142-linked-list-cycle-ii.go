package algorithms

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	x1, x2, x3 := head, head.Next, head.Next.Next
	for x2 != x3 {
		if x3 == nil || x3.Next == nil || x3.Next.Next == nil {
			return nil
		}
		x1 = x1.Next
		x2 = x2.Next.Next
		x3 = x3.Next.Next.Next
	}
	st := head
	x1 = x1.Next
	for st != x1 {
		st = st.Next
		x1 = x1.Next
	}
	return st
}
