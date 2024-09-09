package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
给你一个链表的头节点 head ，该链表包含由 0 分隔开的一连串整数。链表的 开端 和 末尾 的节点都满足 Node.val == 0 。

对于每两个相邻的 0 ，请你将它们之间的所有节点合并成一个节点，其值是所有已合并节点的值之和。然后将所有 0 移除，修改后的链表不应该含有任何 0 。

 返回修改后链表的头节点 head 。
*/

func mergeNodes(head *ListNode) *ListNode {

	dummy := new(ListNode)

	cur := dummy

	pre := 0
	for head != nil {

		if head.Val == 0 {
			if pre != 0 {
				cur.Next = new(ListNode)
				cur = cur.Next
				cur.Val = pre
			}
			pre = 0
		} else {
			pre += head.Val
		}
		head = head.Next
	}
	return dummy.Next
}
