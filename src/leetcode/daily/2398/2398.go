package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
*
你有 n 个机器人，给你两个下标从 0 开始的整数数组 chargeTimes 和 runningCosts ，两者长度都为 n 。
第 i 个机器人充电时间为 chargeTimes[i] 单位时间，花费 runningCosts[i] 单位时间运行。再给你一个整数 budget 。

运行 k 个机器人 总开销 是 max(chargeTimes) + k * sum(runningCosts) ，
其中 max(chargeTimes) 是这 k 个机器人中最大充电时间，sum(runningCosts) 是这 k 个机器人的运行时间之和。

请你返回在 不超过 budget 的前提下，你 最多 可以 连续 运行的机器人数目为多少。

示例 1：

输入：chargeTimes = [3,6,1,3,4], runningCosts = [2,1,3,4,5], budget = 25
输出：3
解释：
可以在 budget 以内运行所有单个机器人或者连续运行 2 个机器人。
选择前 3 个机器人，可以得到答案最大值 3 。总开销是 max(3,6,1) + 3 * sum(2,1,3) = 6 + 3 * 6 = 24 ，小于 25 。
可以看出无法在 budget 以内连续运行超过 3 个机器人，所以我们返回 3 。
示例 2：

输入：chargeTimes = [11,12,19], runningCosts = [10,8,7], budget = 19
输出：0
解释：即使运行任何一个单个机器人，还是会超出 budget，所以我们返回 0 。

提示：

chargeTimes.length == runningCosts.length == n
1 <= n <= 5 * 104
1 <= chargeTimes[i], runningCosts[i] <= 105
1 <= budget <= 1015
*/

type pair struct {
	x int
	y int
}

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {

	n := len(chargeTimes)

	pairs := make([]pair, n)

	for i, v := range chargeTimes {
		pairs[i] = pair{v, runningCosts[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].y < pairs[j].y || pairs[i].y == pairs[j].y && pairs[i].x < pairs[j].x
	})

	hp := new(myHeap)

	s := 0
	cnt := 0
	ans := 0
	for i, v := range pairs {
		heap.Push(hp, pair{v.x, i})
		cnt++
		s += v.y
		for hp.Len() > 0 && int64(cnt*s+(*hp)[0].x) > budget {
			pop := heap.Pop(hp).(pair)
			s -= pairs[pop.y].y
			cnt--
		}
		ans = max(ans, cnt)
	}
	return ans
}

type myHeap []pair

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i].x > (*h)[j].x
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v any) {
	*h = append(*h, v.(pair))
}

func maximumRobots2(chargeTimes []int, runningCosts []int, budget int64) int {

	var st []int
	s := 0
	ans := 0

	l := 0
	for i, v := range runningCosts {
		s += v
		for len(st) > 0 && chargeTimes[st[len(st)-1]] <= chargeTimes[i] {
			st = st[:len(st)-1]
		}
		st = append(st, i)

		for len(st) > 0 && int64(chargeTimes[st[0]]+(i-l+1)*s) > budget {
			s -= runningCosts[l]
			l++
			for len(st) > 0 && st[0] < l {
				st = st[1:]
			}
		}
		ans = max(ans, i-l+1)
	}
	return ans
}

func main() {
	fmt.Println(maximumRobots2([]int{19, 63, 21, 8, 5, 46, 56, 45, 54, 30, 92, 63, 31, 71, 87, 94, 67, 8, 19, 89, 79, 25}, []int{91, 92, 39, 89, 62, 81, 33, 99, 28, 99, 86, 19, 5, 6, 19, 94, 65, 86, 17, 10, 8, 42}, 85))
	fmt.Println(maximumRobots2([]int{32, 83, 96, 70, 98, 80, 30, 42, 63, 67, 49, 10, 80, 13, 69, 91, 73, 10}, []int{49, 67, 92, 26, 18, 22, 38, 34, 36, 26, 32, 84, 39, 42, 88, 51, 8, 2}, 599))
	fmt.Println(maximumRobots([]int{3, 6, 1, 3, 4}, []int{2, 1, 3, 4, 5}, 25))
	fmt.Println(maximumRobots([]int{11, 12, 74, 67, 37, 87, 42, 34, 18, 90, 36, 28, 34, 20}, []int{18, 98, 2, 84, 7, 57, 54, 65, 59, 91, 7, 23, 94, 20}, 937))
}
