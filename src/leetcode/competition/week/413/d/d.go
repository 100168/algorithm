package main

import "fmt"

/*
*
给你一个由 n 个整数组成的数组 nums，以及一个大小为 q 的二维整数数组 queries，其中 queries[i] = [li, ri]。

对于每一个查询，你需要找出 nums[li..ri] 中任意
子数组的 最大异或值。

数组的异或值 需要对数组 a 反复执行以下操作，直到只剩一个元素，剩下的那个元素就是 异或值：

对于除最后一个下标以外的所有下标 i，同时将 a[i] 替换为 a[i] XOR a[i + 1] 。
移除数组的最后一个元素。
返回一个大小为 q 的数组 answer，其中 answer[i] 表示查询 i 的答案。

a b c d                  a b c

a^b b^c c^d           	 a^b b^c

a^c    b^d               a^c

a^c^b^d

示例 1：

输入： nums = [2,8,4,32,16,1], queries = [[0,2],[1,4],[0,5]]

输出： [12,60,60]

f[i][j] = f[i+1][j],f[i][j-1]
*/
func maximumSubarrayXor(nums []int, queries [][]int) []int {

	n := len(nums)
	f := make([][]int, n)
	mx := make([][]int, n)

	for i := range f {
		f[i] = make([]int, n)
		mx[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		f[i][i] = nums[i]
		mx[i][i] = nums[i]
		for j := i + 1; j < n; j++ {
			f[i][j] = f[i+1][j] ^ f[i][j-1]
			mx[i][j] = max(mx[i+1][j], mx[i][j-1], f[i][j])
		}
	}
	ans := make([]int, len(queries))

	for i, v := range queries {
		ans[i] = mx[v[0]][v[1]]
	}
	return ans

}

func main() {
	fmt.Println(maximumSubarrayXor([]int{2, 8, 4, 32, 16, 1}, [][]int{{0, 2}, {1, 4}, {0, 5}}))
}
