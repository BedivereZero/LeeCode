package algorithms

type ListNode struct {
	 Val int
	 Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	root := &ListNode{Next: head}
	curr := root
	for curr.Next != nil {
		next := curr.Next
		count := 1
		for next.Next != nil && curr.Next.Val == next.Next.Val {
			next = next.Next
			count++
		}
		if count > 1 {
			curr.Next = next.Next
		} else {
			curr = curr.Next
		}
	}
	return root.Next
}
