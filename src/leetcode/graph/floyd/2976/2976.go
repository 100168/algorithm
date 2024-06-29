package main

import (
	"fmt"
	"math"
)

/*
*
给你两个下标从 0 开始的字符串 source 和 target ，它们的长度均为 n 并且由 小写 英文字母组成。

另给你两个下标从 0 开始的字符数组 original 和 changed ，以及一个整数数组 cost ，其中 cost[i] 代表将字符 original[i] 更改为字符 changed[i] 的成本。

你从字符串 source 开始。在一次操作中，如果 存在 任意 下标 j 满足 cost[j] == z  、original[j] == x 以及 changed[j] == y 。
你就可以选择字符串中的一个字符 x 并以 z 的成本将其更改为字符 y 。

返回将字符串 source 转换为字符串 target 所需的 最小 成本。如果不可能完成转换，则返回 -1 。

注意，可能存在下标 i 、j 使得 original[j] == original[i] 且 changed[j] == changed[i] 。
*/
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {

	n := len(original)

	w := make([][]int, 26)
	for i := range w {
		w[i] = make([]int, 26)

		for j := range w[i] {
			w[i][j] = math.MaxInt / 2
		}
		w[i][i] = 0
	}
	for i := 0; i < n; i++ {
		f, t := original[i]-'a', changed[i]-'a'
		w[f][t] = min(cost[i], w[f][t])
	}

	f := w

	//f[k][i][j] = min(f[k-1][i][j],f[k-1][i][k]+f[k-1][k][j])
	//f[k+1][i][j] = min(f[k][i][j],f[k][i][k],f[k][k][j])
	//f[k][i][k] == f[k+1][i][k]
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j])
			}
		}
	}

	ans := 0

	for i := range source {
		x, y := source[i]-'a', target[i]-'a'
		if f[x][y] == math.MaxInt/2 {
			return -1
		}
		ans += f[x][y]
	}
	return int64(ans)
}

func main() {

	fmt.Println(minimumCost("abcd", "acbe", []byte("abcced"), []byte("bcbebe"), []int{2, 5, 5, 1, 2, 20}))
}
