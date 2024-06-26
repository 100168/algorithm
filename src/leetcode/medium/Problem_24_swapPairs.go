package main

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
func swapPairs(head *ListNode) *ListNode {
	node := ListNode{Val: 0, Next: head}

	if head == nil || head.Next == nil {
		return head
	}
	current := head.Next
	preNode := head
	parent := &node

	for current != nil {
		next := current.Next
		parent.Next = current
		preNode.Next = next
		current.Next = preNode

		parent = preNode
		if next != nil {
			preNode = next
			current = next.Next
		} else {
			break
		}
	}

	return node.Next
}
func main() {

}
