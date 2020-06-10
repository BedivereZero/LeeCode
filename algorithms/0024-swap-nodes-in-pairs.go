package algorithms

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
	 var second *ListNode
	 if head == nil || head.Next == nil {
		 return head
	 } else {
		 second = head.Next
	 }
	 a, b := head, head.Next
	 for a != nil && b != nil {
		 var c, d *ListNode
		 if b != nil {
			 c = b.Next
		 }
		 if c != nil {
			 d = c.Next
		 }
		 if d == nil {
			 a.Next = c
		 } else {
			 a.Next = d
		 }
		 b.Next = a
		 a, b = c, d
	 }
	 return second
}
