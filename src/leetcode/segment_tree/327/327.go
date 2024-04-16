package main

import "sort"

//给你一个整数数组 nums 以及两个整数 lower 和 upper 。求数组中，值位于范围 [lower, upper] （包含 lower 和
//upper）之内的 区间和的个数 。
//
// 区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。
//
//
//示例 1：
//
//
//输入：nums = [-2,5,-1], lower = -2, upper = 2
//输出：3
//解释：存在三个区间：[0,0]、[2,2] 和 [0,2] ，对应的区间和分别是：-2 、-1 、2 。
//
//
// 示例 2：
//
//
//输入：nums = [0], lower = 0, upper = 0
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
// -10⁵ <= lower <= upper <= 10⁵
// 题目数据保证答案是一个 32 位 的整数
//
//
// Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 582 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

//  s-x>=lower  ==>x<=s-lower
// s-x<=upper  ==>x>=s-upper

type fenwick struct {
	tree []int
}

func (f fenwick) inc(i int) {
	for ; i < len(f.tree); i += i & -i {
		f.tree[i]++
	}
}

func (f fenwick) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}

func (f fenwick) query(l, r int) (res int) {
	return f.sum(r) - f.sum(l-1)
}

func countRangeSum(nums []int, lower, upper int) (cnt int) {
	n := len(nums)

	// 计算前缀和 preSum，以及后面统计时会用到的所有数字 allNums
	allNums := make([]int, 1, 3*n+1)
	preSum := make([]int, n+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
		allNums = append(allNums, preSum[i+1], preSum[i+1]-lower, preSum[i+1]-upper)
	}

	// 将 allNums 离散化
	sort.Ints(allNums)
	k := 1
	kth := map[int]int{allNums[0]: k}
	for i := 1; i <= 3*n; i++ {
		if allNums[i] != allNums[i-1] {
			k++
			kth[allNums[i]] = k
		}
	}

	// 遍历 preSum，利用树状数组计算每个前缀和对应的合法区间数
	t := fenwick{make([]int, k+1)}
	t.inc(kth[0])
	for _, sum := range preSum[1:] {
		left, right := kth[sum-upper], kth[sum-lower]
		cnt += t.query(left, right)
		t.inc(kth[sum])
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
