package main

/**
给你一个整数数组 nums 和一个链表的头节点 head。从链表中移除所有存在于 nums 中的节点后，返回修改后的链表的头节点。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {

	dummy := &ListNode{0, head}

	numsMap := make(map[int]bool)
	for _, v := range nums {
		numsMap[v] = true
	}
	cur := dummy.Next

	pre := dummy

	for cur != nil {
		if numsMap[cur.Val] {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
