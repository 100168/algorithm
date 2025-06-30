package main

import (
	"fmt"
	"strconv"
)

/**
给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。



示例 1：

输入：n = 3, k = 3
输出："213"
示例 2：

输入：n = 4, k = 9
输出："2314"
示例 3：

输入：n = 3, k = 1
输出："123"


提示：

1 <= n <= 9
1 <= k <= n!

143258679

198765432
298765431
*/

func getPermutation(n int, k int) string {

	ans := ""
	f := make([]int, n+1)
	f[0] = 1

	//用来快速计算剩余数的组合数
	for i := 1; i <= n; i++ {
		f[i] = f[i-1] * i
	}

	//用来表示选过的集合
	mask := 0

	//遍历每个位置
	for i := 1; i <= n; i++ {
		//比当前位置小的数的个数
		cnt := 0
		for j := 1; j <= n; j++ {
			//判断当前数是否有没有选
			if mask&(1<<(j-1)) != 0 {
				continue
			}
			cnt++
			//j表示当前数字，f[n-i]表示剩下数的阶乘也就是剩下数的组合数，也就是以j开头可以组合成多少个数
			if cnt*f[n-i] >= k {
				//将当前数加入集合表示已选
				mask |= 1 << (j - 1)
				k -= (cnt - 1) * f[n-i]
				ans += strconv.Itoa(j)
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(getPermutation(9, 33453))
}
