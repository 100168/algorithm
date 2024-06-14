package main

import (
	"fmt"
)

/*
给你有一个 非负 整数 k 。有一个无限长度的台阶，最低 一层编号为 0 。
虎老师有一个整数 jump ，一开始值为 0 。虎老师从台阶 1 开始，虎老师可以使用 任意 次操作，
目标是到达第 k 级台阶。假设虎老师位于台阶 i ，一次 操作 中，虎老师可以：
向下走一级到 i - 1 ，但该操作 不能 连续使用，如果在台阶第 0 级也不能使用。
向上走到台阶 i + 2jump 处，然后 jump 变为 jump + 1 。
请你返回虎老师到达台阶 k 处的总方案数。
注意 ，虎老师可能到达台阶 k 处后，通过一些操作重新回到台阶 k 处，这视为不同的方案。
*/

const mx = 31

var comb [mx][mx]int

func init() {

	comb[0][0] = 1
	for i := 1; i < mx; i++ {
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
		}
	}

}

func waysToReachStair(k int) (ans int) {
	for j := 0; j < mx; j++ {
		m := 1<<j - k
		if m >= 0 && m <= j+1 {
			ans += comb[j+1][m]
		}
	}

	return
}

func main() {
	fmt.Println(waysToReachStair(1))
}
