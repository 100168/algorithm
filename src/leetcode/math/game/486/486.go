package main

/**

给你一个整数数组 nums 。玩家 1 和玩家 2 基于这个数组设计了一个游戏。

玩家 1 和玩家 2 轮流进行自己的回合，玩家 1 先手。开始时，两个玩家的初始分值都是 0 。
每一回合，玩家从数组的任意一端取一个数字（即，nums[0] 或 nums[nums.length - 1]），
取到的数字将会从数组中移除（数组长度减 1 ）。玩家选中的数字将会加到他的得分上。当数组中没有剩余数字可取时，游戏结束。

如果玩家 1 能成为赢家，返回 true 。如果两个玩家得分相等，同样认为玩家 1 是游戏的赢家，也返回 true 。
你可以假设每个玩家的玩法都会使他的分数最大化。

*/

func predictTheWinner(nums []int) bool {
	n := len(nums)

	s := 0

	for i := range nums {
		s += nums[i]
	}
	f1 := make([][]int, n)
	f2 := make([][]int, n)

	for i := range f1 {
		f1[i] = make([]int, n)
		f2[i] = make([]int, n)

		for j := 0; j < n; j++ {
			f1[i][j] = -1
			f2[i][j] = -1
		}
	}

	var second func(int, int) int

	var first func(int, int) int
	second = func(l, r int) int {

		if l == r {
			return 0
		}
		if f2[l][r] != -1 {
			return f2[l][r]
		}
		cur := min(first(l+1, r), first(l, r-1))
		f2[l][r] = cur
		return cur
	}
	first = func(l, r int) int {
		if l == r {
			return nums[l]
		}
		if f1[l][r] != -1 {
			return f1[l][r]
		}
		cur := max(nums[l]+second(l+1, r), nums[r]+second(l, r-1))
		f1[l][r] = cur
		return cur
	}

	l := first(0, n-1)
	return l >= s-l

}
