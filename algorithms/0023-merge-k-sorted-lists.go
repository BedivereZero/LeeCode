package algorithms

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return merge(lists[0], lists[1])
	default:
		return merge(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
	}
}

func merge(a, b *ListNode) *ListNode {
	root := &ListNode{}
	cursor := root
	for a != nil || b != nil {
		if a == nil {
			cursor.Next = b
			b = b.Next
		} else if b == nil {
			cursor.Next = a
			a = a.Next
		} else if a.Val < b.Val {
			cursor.Next = a
			a = a.Next
		} else {
			cursor.Next = b
			b = b.Next
		}
		cursor = cursor.Next
	}
	return root.Next
}
