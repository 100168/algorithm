package main

import (
	"math/rand"
	"slices"
	"sort"
	"strconv"
)

/*
*
给出一个含有不重复整数元素的数组 arr ，每个整数 arr[i] 均大于 1。

用这些整数来构建二叉树，每个整数可以使用任意次数。其中：每个非叶结点的值应等于它的两个子结点的值的乘积。

满足条件的二叉树一共有多少个？答案可能很大，返回 对 109 + 7 取余 的结果。

输入: arr = [2, 4]
输出: 3
解释: 可以得到这些二叉树: [2], [4], [4, 2, 2]
*/
func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)

	mod := int(1e9 + 7)
	ans := 0
	cntMap := make(map[int]int)
	for _, v := range arr {
		cntMap[v]++
		ans++
		for k := 0; arr[k]*arr[k] <= v; k++ {
			if v%arr[k] == 0 {
				t := 1
				if v/arr[k] != arr[k] {
					t = 2
				}
				cntMap[v] = cntMap[v] + (cntMap[arr[k]]*cntMap[v/arr[k]]*t)%mod
				ans = (ans + cntMap[arr[k]]*cntMap[v/arr[k]]*t) % mod
			}
		}
	}
	return ans
}

func numGenerator(maxVal int, maxL int) string {

	ans := make([]int, 0)
	l := rand.Intn(maxL)
	for i := 0; i < l; i++ {
		v := rand.Intn(maxVal)
		ans = append(ans, v)
	}
	sort.Ints(ans)
	ans = slices.Compact(ans)
	s := "["
	for _, v := range ans {
		s += strconv.Itoa(v) + ","
	}
	return s + "]"

}
func main() {
	println(numFactoredBinaryTrees([]int{2, 4}))
	println(numFactoredBinaryTrees([]int{2, 4, 5, 10}))
	//fmt.Println(numGenerator(1e9, 1000))

}
