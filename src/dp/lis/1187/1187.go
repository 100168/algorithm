package main

import (
	"math"
	"sort"
)

//给你两个整数数组 arr1 和 arr2，返回使 arr1 严格递增所需要的最小「操作」数（可能为 0）。
//
// 每一步「操作」中，你可以分别从 arr1 和 arr2 中各选出一个索引，分别为 i 和 j，0 <= i < arr1.length 和 0 <= j
//< arr2.length，然后进行赋值运算 arr1[i] = arr2[j]。
//
// 如果无法让 arr1 严格递增，请返回 -1。
//
//
//
// 示例 1：
//
//
//输入：arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]
//输出：1
//解释：用 2 来替换 5，之后 arr1 = [1, 2, 3, 6, 7]。
//
//
// 示例 2：
//
//
//输入：arr1 = [1,5,3,6,7], arr2 = [4,3,1]
//输出：2
//解释：用 3 来替换 5，然后用 4 来替换 3，得到 arr1 = [1, 3, 4, 6, 7]。
//
//
// 示例 3：
//
//
//输入：arr1 = [1,5,3,6,7], arr2 = [1,6,3,3]
//输出：-1
//解释：无法使 arr1 严格递增。
//
//
//
// 提示：
//
//
// 1 <= arr1.length, arr2.length <= 2000
// 0 <= arr1[i], arr2[i] <= 10^9
//
//
//
// Related Topics 数组 二分查找 动态规划 排序 👍 215 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	m := len(arr1)
	sort.Ints(arr2)

	var dfs func(int, int) int

	cache := make([]map[int]int, m)

	for i := range cache {
		cache[i] = make(map[int]int)
	}
	dfs = func(i, pre int) int {
		if i < 0 {
			return 0
		}
		if _, ok := cache[i][pre]; ok {
			return cache[i][pre]
		}
		cur := math.MaxInt / 2
		minIndex := search(arr2, pre)
		if minIndex >= 0 {
			cur = min(dfs(i-1, arr2[minIndex])+1, cur)
		}
		if arr1[i] < pre {
			cur = min(cur, dfs(i-1, arr1[i]))
		}
		cache[i][pre] = cur
		return cur
	}
	ans := dfs(m-1, math.MaxInt)
	if ans > m {
		return -1
	}
	return ans
}

func search(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) >> 1
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func main() {
	println(makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 4, 3}))
	println(makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 6, 3, 3}))
}

//leetcode submit region end(Prohibit modification and deletion)
