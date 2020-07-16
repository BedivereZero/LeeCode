package main

// Node is a node
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ListNode is a list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
