package main

import (
	"fmt"
	"slices"
	"sort"
)

/**
你有两个果篮，每个果篮中有 n 个水果。给你两个下标从 0 开始的整数数组 basket1 和 basket2 ，
用以表示两个果篮中每个水果的交换成本。你想要让两个果篮相等。为此，可以根据需要多次执行下述操作：

选中两个下标 i 和 j ，并交换 basket1 中的第 i 个水果和 basket2 中的第 j 个水果。
交换的成本是 min(basket1i,basket2j) 。
根据果篮中水果的成本进行排序，如果排序后结果完全相同，则认为两个果篮相等。

返回使两个果篮相等的最小交换成本，如果无法使两个果篮相等，则返回 -1 。


示例 1：

输入：basket1 = [4,2,2,2], basket2 = [1,4,1,2]
输出：1
解释：交换 basket1 中下标为 1 的水果和 basket2 中下标为 0 的水果，交换的成本为 1 。此时，basket1 = [4,1,2,2] 且 basket2 = [2,4,1,2] 。重排两个数组，发现二者相等。
示例 2：

输入：basket1 = [2,3,4,1], basket2 = [3,2,5,1]
输出：-1
解释：可以证明无法使两个果篮相等。


*/

func minCost(basket1 []int, basket2 []int) int64 {

	all := make([]int, len(basket1)*2)
	basket1Map := make(map[int]int)
	basket2Map := make(map[int]int)
	allMap := make(map[int]int)
	for i := 0; i < len(basket1); i++ {
		all[i*2] = basket1[i]
		all[i*2+1] = basket2[i]
		basket1Map[basket1[i]]++
		basket2Map[basket2[i]]++
		allMap[basket1[i]]++
		allMap[basket2[i]]++

	}

	for k := range allMap {

		if allMap[k]&1 != 0 {
			return -1
		}
	}
	sort.Ints(all)

	diff1 := make([]int, 0)
	diff2 := make([]int, 0)

	for i := 0; i < len(all); i += 2 {
		if basket1Map[all[i]]--; basket1Map[all[i]] < 0 {
			diff1 = append(diff1, all[i])
		}
		if basket2Map[all[i]]--; basket2Map[all[i]] < 0 {
			diff2 = append(diff2, all[i])
		}
	}
	minV := all[0]

	slices.Reverse(diff2)
	ans := 0
	for i := range diff1 {
		ans += min(minV*2, min(diff2[i], diff1[i]))
	}

	return int64(ans)

}

// [8 14 43 43 80 80 84 88 88 100]
// [8 14 32 32 42 42 68 68 84 100]

// 32 42 68
// 43 80 88
func main() {
	fmt.Println(minCost([]int{3, 5}, []int{4, 2}))
}
