package main

/*
*
给你一个整数数组 hours，表示以 小时 为单位的时间，返回一个整数，表示满足 i < j 且 hours[i] + hours[j] 构成 整天 的下标对 i, j 的数目。

整天 定义为时间持续时间是 24 小时的 整数倍 。

例如，1 天是 24 小时，2 天是 48 小时，3 天是 72 小时，以此类推。
*/
func countCompleteDayPairs(hours []int) int64 {

	randMap := make(map[int]int)

	for _, h := range hours {
		randMap[h%24]++
	}
	cnt := 0
	for i := 1; i < 12; i++ {
		cnt += randMap[i] * randMap[24-i]
	}

	cnt += randMap[0] * (randMap[0] - 1) / 2
	cnt += randMap[12] * (randMap[12] - 1) / 2
	return int64(cnt)
}
