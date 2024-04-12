package main

import "fmt"

//Range模块是跟踪数字范围的模块。设计一个数据结构来跟踪表示为 半开区间 的范围并查询它们。
//
// 半开区间 [left, right) 表示所有 left <= x < right 的实数 x 。
//
// 实现 RangeModule 类:
//
//
// RangeModule() 初始化数据结构的对象。
// void addRange(int left, int right) 添加 半开区间 [left, right)，跟踪该区间中的每个实数。添加与当前跟踪的
//数字部分重叠的区间时，应当添加在区间 [left, right) 中尚未跟踪的任何数字到该区间中。
// boolean queryRange(int left, int right) 只有在当前正在跟踪区间 [left, right) 中的每一个实数时，才返
//回 true ，否则返回 false 。
// void removeRange(int left, int right) 停止跟踪 半开区间 [left, right) 中当前正在跟踪的每个实数。
//
//
//
//
// 示例 1：
//
//
//输入
//["RangeModule", "addRange", "removeRange", "queryRange", "queryRange",
//"queryRange"]
//[[], [10, 20], [14, 16], [10, 14], [13, 15], [16, 17]]
//输出
//[null, null, null, true, false, true]
//
//解释
//RangeModule rangeModule = new RangeModule();
//rangeModule.addRange(10, 20);
//rangeModule.removeRange(14, 16);
//rangeModule.queryRange(10, 14); 返回 true （区间 [10, 14) 中的每个数都正在被跟踪）
//rangeModule.queryRange(13, 15); 返回 false（未跟踪区间 [13, 15) 中像 14, 14.03, 14.17 这样
//的数字）
//rangeModule.queryRange(16, 17); 返回 true （尽管执行了删除操作，区间 [16, 17) 中的数字 16 仍然会被跟踪）
//
//
//
//
//
// 提示：
//
//
// 1 <= left < right <= 10⁹
// 在单个测试用例中，对 addRange 、 queryRange 和 removeRange 的调用总数不超过 10⁴ 次
//
//
// Related Topics 设计 线段树 有序集合 👍 280 👎 0

// RangeModule leetcode submit region begin(Prohibit modification and deletion)
type RangeModule struct {
	t seg
}

func Constructor() RangeModule {
	t := new(seg)
	t.root = newNode(1, 1e9)
	return RangeModule{t: *t}
}

func (r *RangeModule) AddRange(left int, right int) {
	t := r.t
	t.update(t.root, left, right-1, true)
}

func (r *RangeModule) QueryRange(left int, right int) bool {
	t := r.t
	return t.query(t.root, left, right-1)
}

func (r *RangeModule) RemoveRange(left int, right int) {
	t := r.t
	t.update(t.root, left, right-1, false)
}

type node struct {
	l     int
	r     int
	m     int
	val   bool
	add   bool
	left  *node
	right *node
}

type seg struct {
	root *node
}

func (t seg) query(cur *node, l, r int) bool {
	if cur.l >= l && cur.r <= r {
		return cur.val
	}
	ans := true
	pushDown(cur)
	if l <= cur.m {
		ans = t.query(cur.left, l, r)
	}
	if r > cur.m {
		ans = ans && t.query(cur.right, l, r)
	}
	return ans
}

func (t seg) update(cur *node, l, r int, val bool) {
	if cur.l >= l && cur.r <= r {
		cur.val = val
		cur.add = true
		return
	}
	pushDown(cur)
	m := cur.m
	if l <= m {
		t.update(cur.left, l, r, val)
	}
	if r > m {
		t.update(cur.right, l, r, val)
	}
	pushUp(cur)
}

func pushUp(cur *node) {
	cur.val = cur.left.val && cur.right.val
}

func newNode(l, r int) *node {
	cur := new(node)
	cur.l = l
	cur.r = r
	cur.m = (l + r) >> 1
	return cur
}
func pushDown(cur *node) {
	if cur.left == nil {
		cur.left = newNode(cur.l, cur.m)
	}
	if cur.right == nil {
		cur.right = newNode(cur.m+1, cur.r)
	}
	if cur.add {
		cur.add = false
		cur.left.val = cur.val
		cur.right.val = cur.val
		cur.left.add = true
		cur.right.add = true
	}
}

func main() {
	ran := Constructor()
	ran.AddRange(10, 20)
	ran.RemoveRange(14, 16)
	fmt.Println(ran.QueryRange(10, 15))
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
