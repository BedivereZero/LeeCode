package algorithms

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	newHead := &Node{Val: head.Val}
	record := map[*Node]*Node{head: newHead}

	for old, new := head, newHead; old.Next != nil; old, new = old.Next, new.Next {
		n := &Node{Val: old.Next.Val}
		new.Next = n
		record[old.Next] = n
	}

	for old, new := head, newHead; old != nil; old, new = old.Next, new.Next {
		new.Random = record[old.Random]
	}

	return newHead
}
