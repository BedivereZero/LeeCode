package algorithms

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := &Stack{}, &Stack{}
	for cur := headA; cur != nil; {
		a.push(cur)
		cur = cur.Next
	}
	for cur := headB; cur != nil; {
		b.push(cur)
		cur = cur.Next
	}

	var n *ListNode
	for a.depth > 0 && b.depth > 0 {
		an, bn := a.pop(), b.pop()
		if an != bn {
			break
		} else {
			n = an.(*ListNode)
		}
	}
	return n
}
