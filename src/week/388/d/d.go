package main

import "math"

/*给你一个长度为 n 下标从 0 开始的整数数组 nums 和一个 正奇数 整数 k 。

x 个子数组的能量值定义为 strength = sum[1] * x - sum[2] * (x - 1) + sum[3] * (x - 2) - sum[4] * (x - 3) + ... + sum[x] * 1
，其中 sum[i] 是第 i 个子数 组的和。更正式的，能量值是满足 1 <= i <= x 的所有 i 对应的 (-1)i+1 * sum[i] * (x - i + 1) 之和。

你需要在 nums 中选择 k 个 不相交子数组 ，使得 能量值最大 。

请你返回可以得到的 最大能量值 。

注意，选出来的所有子数组 不 需要覆盖整个数组。*/

func maximumStrength(nums []int, k int) int64 {

	n := len(nums)

	cache := make([][]int64, n)
	for i := range cache {
		cache[i] = make([]int64, k+1)
		for j := range cache[i] {
			cache[i][j] = math.MinInt / 2
		}
	}

	var dfs func(int, int, int64, bool) bool
	dfs = func(index, rest int, sum int64, isPreSelect bool) bool {
		if rest < 0 {
			return false
		}
		if index == n {
			if rest == 0 {
				return true
			}
			return false
		}
		flag1 := int64(1)
		if rest%2 == 0 {
			flag1 = int64(-1)
		}
		flag2 := int64(1)
		if (rest-1)%2 == 0 {
			flag2 = int64(-1)
		}
		ans := int64(math.MaxInt / 2)
		//当前位置不选
		if dfs(index+1, rest, sum, false) {
			ans = sum
		}
		//当前位置作为下一个的开始，重新选
		if dfs(index+1, rest-1, sum+int64(nums[index])*flag2, true) {
			ans = max(sum+int64(nums[index])*flag2, ans)
		}

		if isPreSelect {
			if dfs(index+1, rest, sum+int64(nums[index])*flag1, true) {
				ans = max(ans, sum+int64(nums[index])*flag1)
			}

		}
		cache[index][rest] = ans
		return true
	}
	dfs(0, k, 0, false)
	return cache[0][k]
}

func main() {
	println(maximumStrength([]int{12, -2, -2, -2, -2}, 5))
}
