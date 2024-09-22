package main

import (
	"fmt"
)

/*
*
你写下了若干 正整数 ，并将它们连接成了一个字符串 num 。但是你忘记给这些数字之间加逗号了。你只记得这一列数字是 非递减 的且 没有 任何数字有前导 0 。

请你返回有多少种可能的 正整数数组 可以得到字符串 num 。由于答案可能很大，将结果对 109 + 7 取余 后返回。

示例 1：

输入：num = "327"
输出：2
解释：以下为可能的方案：
3, 27
327
示例 2：

输入：num = "094"
输出：0
解释：不能有数字有前导 0 ，且所有数字均为正数。
示例 3：

输入：num = "0"
输出：0
解释：不能有数字有前导 0 ，且所有数字均为正数。
示例 4：

输入：num = "9999999999999"
输出：101

提示：

1 <= num.length <= 3500
num 只含有数字 '0' 到 '9' 。
*/

func numberOfCombinations(num string) int {

	mod := int(1e9 + 7)
	n := len(num)
	f := make([]int, n)

	sum := make([][]int, n+1)

	for i := range sum {
		sum[i] = make([]int, n+1)

		if num[0] != '0' {
			sum[i][1] = 1
		}
	}

	for i := 0; i < n; i++ {
		if num[i] != '0' {
			f[0] = 1
		}
		for j := 1; j <= i; j++ {
			if num[j] == '0' {
				sum[i+1][j+1] = sum[i+1][j]
				continue
			}

			//优化过程
			//for k := j - 1; k >= 0 && num[k:j] <= num[j:i+1]; k-- {
			//	if num[k] == '0' {
			//		continue
			//	}
			//	f[i][j] = (f[i][j] + f[j-1][k]) % mod
			//}

			//注意取值范围
			t := 2*j - i - 1
			if t >= 0 && num[t:j] > num[j:i+1] {
				t++
			}
			f[j] = (sum[j][j] - sum[j][max(t, 0)] + mod) % mod
			sum[i+1][j+1] = (sum[i+1][j] + f[j]) % mod
		}
	}
	return sum[n][n]
}

func main() {
	fmt.Println(numberOfCombinations("9999999999999"))
}
