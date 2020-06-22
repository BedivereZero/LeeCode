package algorithms

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	root := &ListNode{Next: head}
	cur := root
	for i := 0; i < m - 1; i++ {
		cur = cur.Next
	}
	pointM := cur
	for i := m - 1; i < n; i++ {
		cur = cur.Next
	}
	pointN := cur
	pointM.Next = reverse(pointM.Next, pointN.Next)
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

// func generate(s []int) *ListNode {
// 	root := ListNode{}
// 	cur := &root
// 	for _, v := range s {
// 		cur.Next = &ListNode{Val: v}
// 		cur = cur.Next
// 	}
// 	return root.Next
// }

// func display(h *ListNode) {
// 	fmt.Printf("root -> ")
// 	for h != nil {
// 		fmt.Printf("%d -> ", h.Val)
// 		h = h.Next
// 	}
// 	fmt.Printf("nill\n")
// }
