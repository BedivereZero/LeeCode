package algorithms

type ListNode struct {
    Val     int
    Next    *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    length := circle(head)
    tail := getNode(head, length - (k % length) - 1)
    return part(tail)
}

// Circle list, return length
func circle(head *ListNode) int {
    length := 0
    orig := head
    for head.Next != nil {
        head = head.Next
        length++
    }
    head.Next = orig
    return length + 1
}

func getNode(head *ListNode, n int) *ListNode {
    for i := 0; i < n; i++ {
        head = head.Next
    }
    return head
}

// part circle and return new head
func part(node *ListNode) *ListNode {
    head := node.Next
    node.Next = nil
    return head
}
