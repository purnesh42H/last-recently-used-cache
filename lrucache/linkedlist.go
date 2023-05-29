package lrucache

type ListNode struct {
	left  *ListNode
	right *ListNode
	key   string
	value string
}

func DeleteListNode(node *ListNode) {
	left := node.left
	right := node.right
	left.right = right
	right.left = left
}
