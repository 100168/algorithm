package main

import (
	"container/heap"
	"fmt"
)

/**
给你一个整数 power 和两个整数数组 damage 和 health ，两个数组的长度都为 n 。

Bob 有 n 个敌人，如果第 i 个敌人还活着（也就是健康值 health[i] > 0 的时候），每秒钟会对 Bob 造成 damage[i] 点 伤害。

每一秒中，在敌人对 Bob 造成伤害 之后 ，Bob 会选择 一个 还活着的敌人进行攻击，该敌人的健康值减少 power 。

请你返回 Bob 将 所有 n 个敌人都消灭之前，最少 会收到多少伤害。



示例 1：

输入：power = 4, damage = [1,2,3,4], health = [4,5,6,8]

输出：39

解释：

最开始 2 秒内都攻击敌人 3 ，然后敌人 3 会被消灭，这段时间内对 Bob 的总伤害是 10 + 10 = 20 点。
接下来 2 秒内都攻击敌人 2 ，然后敌人 2 会被消灭，这段时间内对 Bob 的总伤害是 6 + 6 = 12 点。
接下来 1 秒内都攻击敌人 0 ，然后敌人 0 会被消灭，这段时间内对 Bob 的总伤害是 3 点。
接下来 2 秒内都攻击敌人 1 ，然后敌人 1 会被消灭，这段时间内对 Bob 的总伤害是 2 + 2 = 4 点。
*/

func minDamage(power int, damage []int, health []int) int64 {

	n := len(damage)

	pairs := make([]pair, n)

	s := 0
	for i := range pairs {
		pairs[i] = pair{damage[i], health[i] / power}
		if health[i]%power != 0 {
			pairs[i].h += 1
		}
		s += pairs[i].d
	}

	h := new(myHeap)

	for _, pair := range pairs {
		heap.Push(h, pair)
	}

	ans := 0

	for h.Len() > 0 {

		p := heap.Pop(h).(pair)

		ans += s * p.h
		s -= p.d

	}

	return int64(ans)
}

type pair struct{ d, h int }

type myHeap []pair

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i].d*(*h)[j].h > (*h)[j].d*(*h)[i].h
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

func main() {
	fmt.Println(minDamage(1, []int{80, 79}, []int{100000000, 100000000}))
}
