/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	a := head
	for i := 0; i < n; i++ {
		a = a.Next
	}
	if a == nil {
		return head.Next
	}
	b := head
	for a != nil {
		a, b = a.Next, b.Next
	}
	b.Next = b.Next.Next
	return head
}
