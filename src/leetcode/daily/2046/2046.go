package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
给你一个链表的头结点 head ，这个链表是根据结点的绝对值进行升序排序, 返回重新根据节点的值进行升序排序的链表。
*/

func sortLinkedList(head *ListNode) *ListNode {

	poDummy := new(ListNode)
	neDummy := new(ListNode)
	po := poDummy
	for head != nil {
		curNext := head.Next
		if head.Val < 0 {
			temp := neDummy.Next
			neDummy.Next = head
			head.Next = temp
		} else {
			po.Next = head
			po = po.Next
			head.Next = nil
		}
		head = curNext
	}
	nePo := neDummy
	for nePo.Next != nil {
		nePo = nePo.Next
	}
	nePo.Next = poDummy.Next

	return neDummy.Next

}
