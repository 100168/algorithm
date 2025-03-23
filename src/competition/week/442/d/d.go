package main

import (
	"fmt"
	"math/bits"
)

// 返回 [1,n] 的单个元素的操作次数之和
func f(n int) (res int) {
	m := bits.Len(uint(n))
	for i := 1; i < m; i++ {
		res += (i + 1) / 2 << (i - 1)
	}
	return res + (m+1)/2*(n+1-1<<m>>1)
}

func minOperations(queries [][]int) int64 {
	ans := 0
	for _, q := range queries {
		ans += (f(q[1]) - f(q[0]-1) + 1) / 2
	}
	return int64(ans)
}

func main() {
	fmt.Println(minOperations([][]int{{1, 2}, {2, 4}}))
}
