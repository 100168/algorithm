package main

import (
	"fmt"
)

/*
*
给你一个正整数数组 nums 。

如果数组 nums 的子集中的元素乘积是一个 无平方因子数 ，则认为该子集是一个 无平方 子集。

无平方因子数 是无法被除 1 之外任何平方数整除的数字。

返回数组 nums 中 无平方 且 非空 的子集数目。因为答案可能很大，返回对 109 + 7 取余的结果。

nums 的 非空子集 是可以由删除 nums 中一些元素（可以不删除，但不能全部删除）得到的一个数组。
如果构成两个子集时选择删除的下标不同，则认为这两个子集不同。
*/
func squareFreeSubsets(nums []int) int {

	cntOdd := 0
	cnt := make(map[int]int)

	for i := range nums {
		if nums[i] != 1 && (nums[i]%4 == 0 || nums[i]%9 == 0 || nums[i]%25 == 0) {
			continue
		}
		cnt[nums[i]]++
		if nums[i]%2 != 0 {
			cntOdd++
		}
	}

	oddNums := make([]int, 0)
	evenNums := make([]int, 0)

	for k := range cnt {
		if k%2 != 0 {
			oddNums = append(oddNums, k)
		} else {
			evenNums = append(evenNums, k)
		}
	}
	n := len(oddNums)
	mod := int(1e9 + 7)
	ans := 0
	for i := range evenNums {
	next:
		for j := 0; j < 1<<n; j++ {
			c := evenNums[i]
			res := cnt[c]
			for k := 0; k < n; k++ {
				if 1<<k&j != 0 {
					if gcd(c, oddNums[k]) != 1 {
						continue next
					}
					c *= oddNums[k]
					res = (res * cnt[oddNums[k]]) % mod
				}
			}
			ans += res
			ans %= mod
		}
	}

next2:
	for j := 1; j < 1<<n; j++ {
		c := 1
		res := 1
		for k := 0; k < n; k++ {
			if 1<<k&j != 0 {
				if gcd(c, oddNums[k]) != 1 {
					continue next2
				}
				c *= oddNums[k]
				res = (res * cnt[oddNums[k]]) % mod
			}
		}
		ans += res
		ans %= mod
	}

	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	fmt.Println(squareFreeSubsets([]int{3, 4, 4, 5}))
	fmt.Println(squareFreeSubsets([]int{1, 2, 2, 11, 11, 11, 5, 6, 3, 3, 30}))
	fmt.Println(squareFreeSubsets([]int{11, 2, 19, 7, 9, 27}))
}

var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
var sf2mask = [31]int{} // sf2mask[i] 为 i 的质因子集合（用二进制表示）

func init() {
	for i := 2; i <= 30; i++ {
		for j, p := range primes {
			if i%p == 0 {
				if i%(p*p) == 0 { // 有平方因子
					sf2mask[i] = -1
					break
				}
				sf2mask[i] |= 1 << j // 把 j 加到集合中
			}
		}
	}
}

func squareFreeSubsets2(a []int) int {
	const mod int = 1e9 + 7
	cnt, pow2 := [31]int{}, 1
	for _, x := range a {
		if x == 1 {
			pow2 = pow2 * 2 % mod
		} else {
			cnt[x]++
		}
	}

	const m = 1 << len(primes)
	f := [m]int{pow2} // f[j] 表示恰好组成质数集合 j 的方案数，其中用 1 组成空质数集合的方案数为 pow2
	for sf, mask := range sf2mask {
		if mask > 0 && cnt[sf] > 0 {
			other := (m - 1) ^ mask // mask 的补集
			for j := other; ; {     // 枚举 other 的子集 j
				f[j|mask] = (f[j|mask] + f[j]*cnt[sf]) % mod // 不选 mask + 选 mask
				j = (j - 1) & other
				if j == other {
					break
				}
			}
		}
	}
	ans := -1 // 去掉空集（nums 的空子集）
	for _, v := range f {
		ans += v
	}
	return ans % mod
}
