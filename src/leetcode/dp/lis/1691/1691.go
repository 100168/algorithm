package main

import (
	"slices"
	"sort"
)

//给你 n 个长方体 cuboids ，其中第 i 个长方体的长宽高表示为 cuboids[i] = [widthi, lengthi, heighti]（下
//标从 0 开始）。请你从 cuboids 选出一个 子集 ，并将它们堆叠起来。
//
// 如果 widthi <= widthj 且 lengthi <= lengthj 且 heighti <= heightj ，你就可以将长方体 i 堆叠在
//长方体 j 上。你可以通过旋转把长方体的长宽高重新排列，以将它放在另一个长方体上。
//
// 返回 堆叠长方体 cuboids 可以得到的 最大高度 。
//
//
//
// 示例 1：
//
//
//
//
//输入：cuboids = [[50,45,20],[95,37,53],[45,23,12]]
//输出：190
//解释：
//第 1 个长方体放在底部，53x37 的一面朝下，高度为 95 。
//第 0 个长方体放在中间，45x20 的一面朝下，高度为 50 。
//第 2 个长方体放在上面，23x12 的一面朝下，高度为 45 。
//总高度是 95 + 50 + 45 = 190 。
//
//
// 示例 2：
//
//
//输入：cuboids = [[38,25,45],[76,35,3]]
//输出：76
//解释：
//无法将任何长方体放在另一个上面。
//选择第 1 个长方体然后旋转它，使 35x3 的一面朝下，其高度为 76 。
//
//
// 示例 3：
//
//
//输入：cuboids = [[7,11,17],[7,17,11],[11,7,17],[11,17,7],[17,7,11],[17,11,7]]
//输出：102
//解释：
//重新排列长方体后，可以看到所有长方体的尺寸都相同。
//你可以把 11x7 的一面朝下，这样它们的高度就是 17 。
//堆叠长方体的最大高度为 6 * 17 = 102 。
//
//
//
//
// 提示：
//
//
// n == cuboids.length
// 1 <= n <= 100
// 1 <= widthi, lengthi, heighti <= 100
//
//
// Related Topics 数组 动态规划 排序 👍 146 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func maxHeight(cuboids [][]int) int {

	m := len(cuboids)
	for i := range cuboids {
		sort.Ints(cuboids[i])
	}
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] == cuboids[j][0] {
			if cuboids[i][1] == cuboids[j][1] {
				return cuboids[i][2] < cuboids[j][2]
			}
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][0] < cuboids[j][0]
	})
	dp := make([]int, m+1)

	for i := 1; i <= m; i++ {
		cur := cuboids[i-1]
		for j := 1; j < i; j++ {
			if cuboids[j-1][1] <= cur[1] && cuboids[j-1][2] <= cur[2] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += cur[2]
	}
	return slices.Max(dp)
}
func main() {
	//[[38,25,45],[76,35,3]]
	println(maxHeight([][]int{{38, 25, 45}, {76, 35, 3}}))
}

//leetcode submit region end(Prohibit modification and deletion)
