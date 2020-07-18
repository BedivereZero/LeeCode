package algorithms

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return head == head.Next
	}
	slow, fast := head, head.Next
	for slow != fast && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow == fast
}
