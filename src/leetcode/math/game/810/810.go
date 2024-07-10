package main

import "fmt"

/*
*
黑板上写着一个非负整数数组 nums[i] 。

Alice 和 Bob 轮流从黑板上擦掉一个数字，Alice 先手。如果擦除一个数字后，剩余的所有数字按位异或运算得出的结果等于 0 的话，当前玩家游戏失败。
另外，如果只剩一个数字，按位异或运算得到它本身；如果无数字剩余，按位异或运算结果为 0。

并且，轮到某个玩家时，如果当前黑板上所有数字按位异或运算结果等于 0 ，这个玩家获胜。

假设两个玩家每步都使用最优解，当且仅当 Alice 获胜时返回 true。

1,1,1,2,2,2,
*/
func xorGame(nums []int) bool {

	n := len(nums)

	s := 0

	for i := 0; i < n; i++ {
		s ^= nums[i]
	}
	if s == 0 {
		return true
	}
	return n%2 == 0
}

func main() {
	fmt.Println(xorGame([]int{1, 1, 2}))
	fmt.Println(xorGame([]int{0, 1}))
	fmt.Println(xorGame([]int{0, 0, 1, 1}))
	fmt.Println(xorGame([]int{105, 33, 53, 20, 9, 50, 16, 9, 4, 94, 66, 83, 30, 57, 11, 76, 15, 68, 51, 74, 79, 62, 109, 53, 50, 46, 66, 127, 123, 94, 89, 39, 30, 49, 49, 101, 59, 36, 16, 117, 102, 97, 81, 43, 44, 109, 116, 36, 50, 106, 19}))
}
