package main

import (
	"fmt"
	"sort"
)

/**

你的初始 能量 为 power，初始 分数 为 0，只有一包令牌以整数数组 tokens 给出。其中 tokens[i] 是第 i 个令牌的值（下标从 0 开始）。

你的目标是通过有策略地使用这些令牌以 最大化 总 分数。在一次行动中，
你可以用两种方式中的一种来使用一个 未被使用的 令牌（但不是对同一个令牌使用两种方式）：

朝上：如果你当前 至少 有 tokens[i] 点 能量 ，可以使用令牌 i ，失去 tokens[i] 点 能量 ，并得到 1 分 。
朝下：如果你当前至少有 1 分 ，可以使用令牌 i ，获得 tokens[i] 点 能量 ，并失去 1 分 。
在使用 任意 数量的令牌后，返回我们可以得到的最大 分数 。



示例 1：

输入：tokens = [100], power = 50
输出：0
解释：因为你的初始分数为 0，无法使令牌朝下。你也不能使令牌朝上因为你的能量（50）比 tokens[0] 少（100）。


输入：tokens = [100,200,300,400], power = 200
输出：2
解释：按下面顺序使用令牌可以得到 2 分：
1. 令牌 0 (100)正面朝上，能量变为 100 ，分数变为 1
2. 令牌 3 (400)正面朝下，能量变为 500 ，分数变为 0
3. 令牌 1 (200)正面朝上，能量变为 300 ，分数变为 1
4. 令牌 2 (300)正面朝上，能量变为 0 ，分数变为 2
*/

func bagOfTokensScore(tokens []int, power int) int {

	sort.Ints(tokens)

	n := len(tokens)

	cur := 0
	l, r := 0, n-1

	ans := 0

	for l <= r {
		if tokens[l] <= power {
			power -= tokens[l]
			cur++
			l++
		} else {
			power += tokens[r]
			cur--
			r--
		}
		if cur < 0 {
			break
		}

		ans = max(ans, cur)
	}
	return ans
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
	fmt.Println(bagOfTokensScore([]int{71, 55, 82}, 54))
}
