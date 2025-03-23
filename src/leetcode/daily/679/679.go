package main

import (
	"fmt"
	"math"
)

/*
给定一个长度为4的整数数组 cards 。你有 4 张卡片，每张卡片上都包含一个范围在 [1,9] 的数字。您应该使用运算符 ['+', '-', '*', '/'] 和括号 '(' 和 ')' 将这些卡片上的数字排列成数学表达式，以获得值24。

你须遵守以下规则:

除法运算符 '/' 表示实数除法，而不是整数除法。
例如， 4 /(1 - 2 / 3)= 4 /(1 / 3)= 12 。
每个运算都在两个数字之间。特别是，不能使用 “-” 作为一元运算符。
例如，如果 cards =[1,1,1,1] ，则表达式 “-1 -1 -1 -1” 是 不允许 的。
你不能把数字串在一起
例如，如果 cards =[1,2,1,2] ，则表达式 “12 + 12” 无效。
如果可以得到这样的表达式，其计算结果为 24 ，则返回 true ，否则返回 false 。

示例 1:

输入: cards = [4, 1, 8, 7]
输出: true
解释: (8-4) * (7-1) = 24
示例 2:

输入: cards = [1, 2, 1, 2]
输出: false
*
*/

const (
	EPSILON = 1e-6
)

func judgePoint24(cards []int) bool {

	var dfs func([]float64) bool

	dfs = func(arr []float64) bool {

		if len(arr) == 0 {
			return false
		}

		if len(arr) == 1 {

			return math.Abs(arr[0]-float64(24)) <= EPSILON
		}

		for i := 0; i < len(arr); i++ {

			for j := 0; j < len(arr); j++ {

				if i == j {
					continue
				}

				np := make([]float64, 0)

				for k := 0; k < len(arr); k++ {

					if k != i && k != j {
						np = append(np, arr[k])
					}
				}

				for k := 0; k < 4; k++ {

					switch k {

					case 0:
						np = append(np, arr[i]+arr[j])

					case 1:
						np = append(np, arr[i]-arr[j])
					case 2:

						if arr[j] < EPSILON {
							continue
						}
						np = append(np, arr[i]/arr[j])

					case 3:

						np = append(np, arr[i]*arr[j])

					}

					if dfs(np) {
						return true
					}

					np = np[:len(np)-1]

				}
			}

		}

		return false
	}

	npp := make([]float64, 4)
	for i := 0; i < 4; i++ {
		npp[i] = float64(cards[i])
	}
	return dfs(npp)
}
func main() {
	fmt.Println(judgePoint24([]int{4, 1, 8, 7}))
}
