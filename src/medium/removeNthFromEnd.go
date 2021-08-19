package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	//防止删除头节点，新建一个节点，指向头节点
	dump := &ListNode{0, head}
	fast := head
	slow := dump

	for i := 1; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dump.Next
}
func main() {
	fmt.Println("hh")
}
