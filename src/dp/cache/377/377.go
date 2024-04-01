package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if v > i {
				continue
			}
			dp[i] += dp[i-v]
		}
	}

	return dp[target]

}
func main() {
	println(combinationSum4([]int{1, 2, 3}, 4))
}
