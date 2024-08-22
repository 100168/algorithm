package main

import (
	"fmt"
	"math"
)

/**
给你一组带编号的 balls 并要求将它们分类到盒子里，以便均衡地分配。你必须遵守两条规则：

同一个盒子里的球必须具有相同的编号。但是，如果你有多个相同编号的球，你可以把它们放在不同的盒子里。
最大的盒子只能比最小的盒子多一个球。
返回遵循上述规则排列这些球所需要的盒子的最小数目。



示例 1：

输入：balls = [3,2,3,2,3]
输出：2
解释：一个得到 2 个分组的方案如下，中括号内的数字都是下标：
我们可以如下排列 balls 到盒子里：
- [3,3,3]
- [2,2]
两个盒子之间的大小差没有超过 1。

思路:贪心
1.先统计每个种类出现次数
2.从最小的出现次数从大往小枚举，第一个满足的则是花费最少得
3.
*/

func minGroupsForValidAssignment(balls []int) int {

	cnt := make(map[int]int)

	for _, v := range balls {
		cnt[v]++
	}

	start := math.MaxInt
	for _, v := range cnt {
		start = min(start, v)
	}
next:
	//枚举每个桶最小次数
	for i := start; i > 0; i-- {
		cur := 0
		for _, v := range cnt {
			upSum := v / (i + 1)
			upRest := v % (i + 1)

			//判断是否满足条件
			if upRest > 0 && upSum+upRest < i {
				continue next
			}
			cur += upSum
			if upRest > 0 {
				cur += 1
			}
		}
		return cur
	}
	return 0
}

func main() {
	//fmt.Println(minGroupsForValidAssignment([]int{3, 2, 3, 2, 3}))
	fmt.Println(minGroupsForValidAssignment([]int{100000, 1000001, 1000002, 1000003, 1000004, 1000005}))
}
