package main

import (
	"fmt"
	"slices"
)

//给你一个整数数组 nums 和一个整数 k 。
//
// 找到 nums 中满足以下要求的最长子序列：
//
//
// 子序列 严格递增
// 子序列中相邻元素的差值 不超过 k 。
//
//
// 请你返回满足上述要求的 最长子序列 的长度。
//
// 子序列 是从一个数组中删除部分元素后，剩余元素不改变顺序得到的数组。
//
//
//
// 示例 1：
//
// 输入：nums = [4,2,1,4,3,4,5,8,15], k = 3
//输出：5
//解释：
//满足要求的最长子序列是 [1,3,4,5,8] 。
//子序列长度为 5 ，所以我们返回 5 。
//注意子序列 [1,3,4,5,8,15] 不满足要求，因为 15 - 8 = 7 大于 3 。
//
//
// 示例 2：
//
// 输入：nums = [7,4,5,1,8,12,4,7], k = 5
//输出：4
//解释：
//满足要求的最长子序列是 [4,5,8,12] 。
//子序列长度为 4 ，所以我们返回 4 。
//
//
// 示例 3：
//
// 输入：nums = [1,5], k = 1
//输出：1
//解释：
//满足要求的最长子序列是 [1] 。
//子序列长度为 1 ，所以我们返回 1 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i], k <= 10⁵
//
//
// Related Topics 树状数组 线段树 队列 数组 分治 动态规划 单调队列 👍 77 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

type seg []struct {
	maxVal int
	l      int
	r      int
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}
func (t seg) update(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].maxVal = v
		return
	}
	m := (t[o].l + t[o].r) / 2

	if i <= m {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t[o].maxVal = max(t[o<<1].maxVal, t[o<<1|1].maxVal)
}

func (t seg) query(o, l, r int) int {

	if t[o].l >= l && t[o].r <= r {
		return t[o].maxVal
	}
	m := (t[o].l + t[o].r) / 2

	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l >= m {
		return t.query(o<<1|1, l, r)
	}

	return max(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func lengthOfLIS(nums []int, k int) int {

	mx := slices.Max(nums)
	t := make(seg, mx*4)
	t.build(1, 1, mx)

	for _, x := range nums {
		if x == 1 {
			t.update(1, 1, 1)
		} else {
			t.update(1, x, 1+t.query(1, max(x-k, 1), x-1))
		}
	}
	return t[1].maxVal
}

func main() {
	fmt.Println(lengthOfLIS([]int{1, 5}, 4))
}

//leetcode submit region end(Prohibit modification and deletion)
