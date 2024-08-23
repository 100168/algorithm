package main

import "math/bits"

/*
*
一个非负整数 x 的 强数组 指的是满足元素为 2 的幂且元素总和为 x 的最短有序数组。下表说明了如何确定 强数组 的示例。
可以证明，x 对应的强数组是独一无二的。

数字	二进制表示	强数组
1	00001	[1]
8	01000	[8]
10	01010	[2, 8]
13	01101	[1, 4, 8]
23	10111	[1, 2, 4, 16]

我们将每一个升序的正整数 i （即1，2，3等等）的 强数组 连接得到数组 big_nums ，big_nums 开始部分为 [1, 2, 1, 2, 4, 1, 4, 2, 4, 1, 2, 4, 8, ...] 。

给你一个二维整数数组 queries ，其中 queries[i] = [fromi, toi, modi] ，你需要计算 (big_nums[fromi] * big_nums[fromi + 1] * ... * big_nums[toi]) % modi 。

请你返回一个整数数组 answer ，其中 answer[i] 是第 i 个查询的答案。
*/
func sumE(k int) (res int) {
	var n, cnt1, sumI int
	for i := bits.Len(uint(k+1)) - 1; i > 0; i-- {
		c := cnt1<<i + i<<(i-1) // 新增的幂次个数
		if c <= k {
			k -= c
			res += sumI<<i + i*(i-1)/2<<(i-1)
			sumI += i   // 之前填的 1 的幂次之和
			cnt1++      // 之前填的 1 的个数
			n |= 1 << i // 填 1
		}
	}
	// 最低位单独计算
	if cnt1 <= k {
		k -= cnt1
		res += sumI
		n |= 1 // 最低位填 1
	}
	// 剩余的 k 个幂次，由 n 的低 k 个 1 补充
	for ; k > 0; k-- {
		res += bits.TrailingZeros(uint(n))
		n &= n - 1 // 去掉最低位的 1（置为 0）
	}
	return
}

func findProductsOfElements(queries [][]int64) []int {
	ans := make([]int, len(queries))
	for i, q := range queries {
		er := sumE(int(q[1]) + 1)
		el := sumE(int(q[0]))
		ans[i] = pow(2, er-el, int(q[2]))
	}
	return ans
}

func pow(x, n, mod int) int {
	res := 1 % mod
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
