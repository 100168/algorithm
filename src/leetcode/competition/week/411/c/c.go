package main

import "fmt"

/**

你两个 正整数 n 和 k。

如果整数 x 满足以下全部条件，则该整数是一个 k 回文数：

x 是一个
回文数
。
x 可以被 k 整除。
以字符串形式返回 最大的  n 位 k 回文数。

注意，该整数 不 含前导零。



示例 1：

输入： n = 3, k = 5

输出： "595"

解释：

595 是最大的 3 位 k 回文数。

示例 2：

输入： n = 1, k = 4

输出： "8"

解释：

1 位 k 回文数只有 4 和 8。

示例 3：

输入： n = 5, k = 6

输出： "89898"





提示：

1 <= n <= 10^5
1 <= k <= 9

*/

func largestPalindrome(n, k int) string {

	pow10 := make([]int, n)
	pow10[0] = 1

	ans := make([]byte, n)
	for i := 1; i < n; i++ {
		pow10[i] = pow10[i-1] * 10 % k
	}

	m := (n + 1) / 2

	visit := make([][]bool, m+1)

	for i := range visit {
		visit[i] = make([]bool, k)
	}

	var dfs func(int, int) bool

	dfs = func(i int, j int) bool {
		if i == m {
			return j == 0
		}

		visit[i][j] = true
		for d := 9; d >= 0; d-- {
			j2 := 0
			if n&1 != 0 && i == m-1 {
				j2 = (pow10[i]*d + j) % k
			} else {
				j2 = (j + (pow10[i]+pow10[n-1-i])*d) % k
			}
			if !visit[i+1][j2] && dfs(i+1, j2) {
				c := '0' + byte(d)
				ans[i] = c
				ans[n-i-1] = ans[i]
				return true
			}

		}

		return false
	}

	dfs(0, 0)
	return string(ans)
}

func main() {
	fmt.Println(largestPalindrome(5, 6))
}
