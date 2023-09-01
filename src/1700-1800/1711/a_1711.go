package main

import "fmt"

// 大餐 是指 恰好包含两道不同餐品 的一餐，其美味程度之和等于 2 的幂。
//
// 你可以搭配 任意 两道餐品做一顿大餐。
//
// 给你一个整数数组 deliciousness ，其中 deliciousness[i] 是第 i 道餐品的美味程度，返回你可以用数组中的餐品做出的不同 大
// 餐 的数量。结果需要对 10⁹ + 7 取余。
//
// 注意，只要餐品下标不同，就可以认为是不同的餐品，即便它们的美味程度相同。
func countPairs(deliciousness []int) int {
	nums := make([]int, (1<<20)+1)

	for i := range deliciousness {
		nums[deliciousness[i]]++
	}
	mod := uint64(100_000_000_7)
	ans := uint64(0)

	for _, v := range deliciousness {
		for j := 0; j <= 21; j++ {
			cur := (1 << j) - v
			if cur >= 0 && cur <= 1<<20 {
				if cur == v {
					ans = ans + uint64(nums[cur]) - 1
				} else {
					ans = ans + uint64(nums[cur])
				}

			}
		}
	}
	return int((ans / 2) % mod)
}

func countPairs2(deliciousness []int) (ans int) {
	maxVal := deliciousness[0]
	for _, val := range deliciousness[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	maxSum := maxVal * 2
	cnt := map[int]int{}
	for _, val := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			ans += cnt[sum-val]
		}
		cnt[val]++
	}
	return ans % (1e9 + 7)
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(countPairs(nums))
}
