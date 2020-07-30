package algorithms

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	for cur := head; cur.Next != nil; {
		if cur.Val <= cur.Next.Val {
			cur = cur.Next
			continue
		}
		if cur.Next.Val <= head.Val {
			cur.Next, cur.Next.Next, head = cur.Next.Next, head, cur.Next
			continue
		}
		pos := head
		for ; pos.Next.Val < cur.Next.Val; pos = pos.Next {
		}
		pos.Next, cur.Next.Next, cur.Next = cur.Next, pos.Next, cur.Next.Next
	}

	return head
}
