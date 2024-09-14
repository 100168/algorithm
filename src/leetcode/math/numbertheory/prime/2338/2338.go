package main

import "fmt"

/**
给你两个整数 n 和 maxValue ，用于描述一个 理想数组 。

对于下标从 0 开始、长度为 n 的整数数组 arr ，如果满足以下条件，则认为该数组是一个 理想数组 ：

每个 arr[i] 都是从 1 到 maxValue 范围内的一个值，其中 0 <= i < n 。
每个 arr[i] 都可以被 arr[i - 1] 整除，其中 0 < i < n 。
返回长度为 n 的 不同 理想数组的数目。由于答案可能很大，返回对 109 + 7 取余的结果。



示例 1：

输入：n = 2, maxValue = 5
输出：10
解释：存在以下理想数组：
- 以 1 开头的数组（5 个）：[1,1]、[1,2]、[1,3]、[1,4]、[1,5]
- 以 2 开头的数组（2 个）：[2,2]、[2,4]
- 以 3 开头的数组（1 个）：[3,3]
- 以 4 开头的数组（1 个）：[4,4]
- 以 5 开头的数组（1 个）：[5,5]
共计 5 + 2 + 1 + 1 + 1 = 10 个不同理想数组。


输入：n = 5, maxValue = 3
输出：11
解释：存在以下理想数组：
- 以 1 开头的数组（9 个）：
   - 不含其他不同值（1 个）：[1,1,1,1,1]
   - 含一个不同值 2（4 个）：[1,1,1,1,2], [1,1,1,2,2], [1,1,2,2,2], [1,2,2,2,2]
   - 含一个不同值 3（4 个）：[1,1,1,1,3], [1,1,1,3,3], [1,1,3,3,3], [1,3,3,3,3]
- 以 2 开头的数组（1 个）：[2,2,2,2,2]
- 以 3 开头的数组（1 个）：[3,3,3,3,3]
共计 9 + 1 + 1 = 11 个不同理想数组。


2,6

1,6
2,6
3,6
6,6

参考：
https://leetcode.cn/problems/count-the-number-of-ideal-arrays/solutions/1659088/shu-lun-zu-he-shu-xue-zuo-fa-by-endlessc-iouh/
*/

const mod = 1_000_000_007

var p [10015]int

const k = 10015

func init() {
	p[0] = 1
	for i := 1; i < k; i++ {
		p[i] = p[i-1] * i % mod
	}
}

func comb(n, m int) int {
	return p[n] * pow(p[m]*p[n-m]%mod, mod-2) % mod
}

func pow(a, b int) int {
	ans := 1
	for ; b > 0; b >>= 1 {
		if b&1 != 0 {
			ans = ans * a % mod
		}
		a = a * a % mod
	}
	return ans
}

func idealArrays(n int, maxValue int) int {

	ans := 1
	for i := 2; i <= maxValue; i++ {
		v := i
		cur := 1
		for j := 2; j*j <= v; j++ {
			cnt := 0
			for v%j == 0 {
				v /= j
				cnt++
			}

			if cnt > 0 {
				cur = cur * comb(cnt+n-1, n-1) % mod
			}
			cnt = 0
		}
		if v > 1 {
			cur = cur * comb(n, n-1) % mod
		}
		ans = (ans + cur) % mod
	}
	return ans
}

func main() {
	fmt.Println(idealArrays(2, 6))
	//fmt.Println(idealArrays(5, 3))
}
