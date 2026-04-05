package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**给定一个头结点为 head 的非空单链表，返回链表的中间结点。 如果有两个中间结点，则返回第二个中间结点。**/
func middleNode(head *ListNode) *ListNode {

	var fast *ListNode
	var slow *ListNode
	fast = head
	slow = head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		if fast != nil {
			fast = fast.Next
		} else {
			break
		}
	}
	return slow
}
func main() {

}
