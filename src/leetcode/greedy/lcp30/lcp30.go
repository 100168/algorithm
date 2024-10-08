package main

import (
	"container/heap"
	"fmt"
)

/**
小扣当前位于魔塔游戏第一层，共有 N 个房间，编号为 0 ~ N-1。
每个房间的补血道具/怪物对于血量影响记于数组 nums，其中正数表示道具补血数值，即血量增加对应数值
；负数表示怪物造成伤害值，即血量减少对应数值；0 表示房间对血量无影响。

小扣初始血量为 1，且无上限。假定小扣原计划按房间编号升序访问所有房间补血/打怪，
为保证血量始终为正值，小扣需对房间访问顺序进行调整，每次仅能将一个怪物房间（负数的房间）调整至访问顺序末尾。
请返回小扣最少需要调整几次，才能顺利访问所有房间。若调整顺序也无法访问完全部房间，请返回 -1。

示例 1：

输入：nums = [100,100,100,-250,-60,-140,-50,-50,100,150]

输出：1

解释：初始血量为 1。至少需要将 nums[3] 调整至访问顺序末尾以满足要求。


*/

func magicTower(nums []int) int {

	h := new(myHeap)
	s := 1
	ops := 0

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < 0 {
		return -1
	}
	for len(nums) > 0 {
		v := nums[0]
		nums = nums[1:]
		s += v
		if v < 0 {
			heap.Push(h, v)
		}
		if s <= 0 {
			pop := heap.Pop(h).(int)
			s -= pop
			nums = append(nums, pop)
			ops++

		}
	}
	return ops

}

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
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
	*h = append(*h, v.(int))
}
func main() {
	fmt.Println(magicTower([]int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}))
	fmt.Println(magicTower([]int{-1, -1, 10}))
}
