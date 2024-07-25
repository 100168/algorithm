package main

import (
	"fmt"
)

/*
*
给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，两个数组的大小都为 n ，同时给你一个整数 diff ，统计满足以下条件的 数对 (i, j) ：

0 <= i < j <= n - 1 且
nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff.

nums1[i] - nums2[i]  <=  nums1[j]- nums2[j]+ diff.

s[i]-diff <= s[j]
请你返回满足条件的 数对数目 。

示例 1：

输入：nums1 = [3,2,5], nums2 = [2,2,1], diff = 1
输出：3
解释：
总共有 3 个满足条件的数对：
1. i = 0, j = 1：3 - 2 <= 2 - 2 + 1 。因为 i < j 且 1 <= 1 ，这个数对满足条件。
2. i = 0, j = 2：3 - 5 <= 2 - 1 + 1 。因为 i < j 且 -2 <= 2 ，这个数对满足条件。
3. i = 1, j = 2：2 - 5 <= 2 - 1 + 1 。因为 i < j 且 -3 <= 2 ，这个数对满足条件。
所以，我们返回 3 。
*/
func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {

	n := len(nums1)
	diffs := make([]int, n)
	for i := range nums1 {
		diffs[i] = nums1[i] - nums2[i]
	}
	sg := new(segmentTree)
	sg.root = createNode(-100001, 100001)

	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans += sg.query(sg.root, diffs[i]-diff, 100001)
		sg.update(sg.root, diffs[i], diffs[i], 1)
	}
	return int64(ans)
}

type segmentTree struct {
	root *node
}
type node struct {
	l     int
	r     int
	m     int
	left  *node
	right *node
	val   int
	add   bool
	sum   int
}

func createNode(l, r int) *node {
	cur := new(node)
	cur.l = l
	cur.r = r
	cur.m = (l + r) >> 1
	return cur
}

func (sg segmentTree) query(cur *node, l, r int) int {
	if cur.l >= l && cur.r <= r {
		return cur.val
	}
	sg.pushDown(cur)
	ans := 0
	if l <= cur.m {
		ans = sg.query(cur.left, l, r)
	}
	if r > cur.m {
		ans = ans + sg.query(cur.right, l, r)
	}
	return ans
}

func (sg segmentTree) update(cur *node, l, r, val int) {

	if cur.l >= l && cur.r <= r {
		cur.val += (cur.r - cur.l + 1) * val
		cur.add = true
		cur.sum += val
		return
	}
	sg.pushDown(cur)
	if l <= cur.m {
		sg.update(cur.left, l, r, val)
	}
	if r > cur.m {
		sg.update(cur.right, l, r, val)
	}
	sg.pushUp(cur)
}
func (sg segmentTree) pushUp(cur *node) {
	cur.val = cur.left.val + cur.right.val
}

func (sg segmentTree) pushDown(cur *node) {
	if cur.left == nil {
		cur.left = createNode(cur.l, cur.m)
	}
	if cur.right == nil {
		cur.right = createNode(cur.m+1, cur.r)
	}
	if cur.add {
		cur.add = false
		cur.left.add = true
		cur.left.val += cur.sum
		cur.left.sum += cur.sum
		cur.right.add = true
		cur.right.val += cur.sum
		cur.right.sum += cur.sum
		cur.sum = 0
	}
}

func main() {

	fmt.Println(numberOfPairs([]int{3, 2, 5}, []int{2, 2, 1}, 1))
	fmt.Println(numberOfPairs([]int{3, -1}, []int{-2, 2}, -1))
}
