package main

/*
*
给你一个下标从 0 开始的二维整数数组 flowers ，
其中 flowers[i] = [starti, endi] 表示第 i 朵花的 花期 从 starti 到 endi （都 包含）。
同时给你一个下标从 0 开始大小为 n 的整数数组 people ，people[i] 是第 i 个人来看花的时间。

请你返回一个大小为 n 的整数数组 answer ，其中 answer[i]是第 i 个人到达时在花期内花的 数目 。
*/
func fullBloomFlowers(flowers [][]int, people []int) []int {

	seg := new(segmentTree)
	maxVal := 0
	for i := range flowers {
		maxVal = max(maxVal, flowers[i][1])
	}
	seg.root = createNode(0, 1<<31)
	for _, v := range flowers {
		seg.update(seg.root, v[0], v[1], 1)
	}
	ans := make([]int, len(people))
	for i, v := range people {
		ans[i] = seg.query(seg.root, v, v)
	}
	return ans
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
