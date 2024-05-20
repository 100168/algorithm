package main

import (
	"fmt"
	"slices"
)

/*
*
给你一个下标从 0 开始的整数数组 nums ，数组长度为 n 。

你可以执行无限次下述运算：

选择一个之前未选过的下标 i ，并选择一个 严格小于 nums[i] 的质数 p ，从 nums[i] 中减去 p 。
如果你能通过上述运算使得 nums 成为严格递增数组，则返回 true ；否则返回 false 。

严格递增数组 中的每个元素都严格大于其前面的元素。

输入：nums = [4,9,6,10]
输出：true
解释：
在第一次运算中：选择 i = 0 和 p = 3 ，然后从 nums[0] 减去 3 ，nums 变为 [1,9,6,10] 。
在第二次运算中：选择 i = 1 和 p = 7 ，然后从 nums[1] 减去 7 ，nums 变为 [1,2,6,10] 。
第二次运算后，nums 按严格递增顺序排序，因此答案为 true 。
*/
func primeSubOperation(nums []int) bool {

	n := slices.Max(nums)
	isPrime := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	prime := make([]int, 0)
	for i := 2; i < len(isPrime); i++ {
		if isPrime[i] {
			prime = append(prime, i)
		}
	}
	pre := 0

	for _, num := range nums {

		//num-pre>x
		x := num - pre
		if x <= 0 {
			return false
		}

		l, r := 0, len(prime)-1

		for l <= r {
			m := (r + l) / 2
			if prime[m] >= x {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		pre = num
		if r != -1 {
			pre -= prime[r]
		}
		println(pre)
	}
	return true
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeSubOperation([]int{4, 9, 6, 10}))
}
