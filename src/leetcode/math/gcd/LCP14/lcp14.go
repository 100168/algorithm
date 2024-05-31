package main

import (
	"fmt"
	"math"
)

/*
*
给定一个整数数组 nums ，小李想将 nums 切割成若干个非空子数组，
使得每个子数组最左边的数和最右边的数的最大公约数大于 1 。
为了减少他的工作量，请求出最少可以切成多少个子数组。

示例 1：

输入：nums = [2,3,3,2,3,3]

输出：2

解释：最优切割为 [2,3,3,2] 和 [3,3] 。第一个子数组头尾数字的最大公约数为 2 ，第二个子数组头尾数字的最大公约数为 3 。

示例 2：

输入：nums = [2,3,5,7]

输出：4

解释：只有一种可行的切割：[2], [3], [5], [7]

限制：

1 <= nums.length <= 10^5

1 <= nums.length <= 10^5

2 <= nums[i] <= 10^6
*/
func splitArray(nums []int) int {

	primes := make(map[int][]int)
	gcdMap := make(map[int][]int)
	for i, v := range nums {
		t := v
		for j := 2; j*j <= v; j++ {
			if v%j == 0 {
				primes[j] = append(primes[j], i)
				gcdMap[t] = append(gcdMap[t], j)
				for v%j == 0 {
					v /= j
				}
			}

		}
		if v > 1 {
			primes[v] = append(primes[v], i)
			gcdMap[t] = append(gcdMap[t], v)
		}

	}

	n := len(nums)

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int

	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		cur := math.MaxInt / 2
		for _, v := range gcdMap[nums[i]] {
			for _, x := range primes[v] {
				if x <= i {
					cur = min(cur, dfs(x-1)+1)
				}

			}
		}
		memo[i] = cur
		return cur
	}
	return dfs(n - 1)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(splitArray([]int{2, 3, 3, 2, 3, 3}))
}
