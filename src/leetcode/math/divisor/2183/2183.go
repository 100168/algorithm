package main

import "fmt"

/*
*
给你一个下标从 0 开始、长度为 n 的整数数组 nums 和一个整数 k ，返回满足下述条件的下标对 (i, j) 的数目：
0 <= i < j <= n - 1 且
nums[i] * nums[j] 能被 k 整除。

输入：nums = [1,2,3,4,5], k = 2
输出：7
解释：
共有 7 对下标的对应积可以被 2 整除：
(0, 1)、(0, 3)、(1, 2)、(1, 3)、(1, 4)、(2, 3) 和 (3, 4)
它们的积分别是 2、4、6、8、10、12 和 20 。
其他下标对，例如 (0, 2) 和 (2, 4) 的乘积分别是 3 和 15 ，都无法被 2 整除。
*/
func countPairs(nums []int, k int) int64 {

	primeKCnt := make(map[int]int)
	for _, n := range nums {
		primeKCnt[gcd(n, k)]++
	}
	ans := 0

	for k1, v1 := range primeKCnt {
		for k2, v2 := range primeKCnt {
			if k1 == k2 {
				v2 -= 1
			}
			if k2*k1%k == 0 {
				ans += v1 * v2
			}
		}
	}

	return int64(ans / 2)
}

func countPairs2(nums []int, k int) (ans int64) {
	divisors := []int{}
	for d := 1; d*d <= k; d++ { // 预处理 k 的所有因子
		if k%d == 0 {
			divisors = append(divisors, d)
			if d*d < k {
				divisors = append(divisors, k/d)
			}
		}
	}
	cnt := map[int]int{}
	for _, v := range nums {
		ans += int64(cnt[k/gcd(v, k)])
		for _, d := range divisors {
			if v%d == 0 {
				cnt[d]++
			}
		}
	}
	return
}

func countPairs3(nums []int, k int) (ans int64) {
	divisors := []int{}
	for d := 1; d*d <= k; d++ { // 预处理 k 的所有因子
		if k%d == 0 {
			divisors = append(divisors, d)
			if d*d < k {
				divisors = append(divisors, k/d)
			}
		}
	}
	cnt := map[int]int{}
	for _, v := range nums {
		ans += int64(cnt[k/gcd(v, k)])
		for _, d := range divisors {
			if v%d == 0 {
				cnt[d]++
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(countPairs([]int{1, 2, 3, 4, 5}, 2))
}
