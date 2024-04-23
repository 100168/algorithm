package main

/*
可以用字符串表示一个学生的出勤记录，其中的每个字符用来标记当天的出勤情况（缺勤、迟到、到场）。记录中只含下面三种字符：
'A'：Absent，缺勤
'L'：Late，迟到
'P'：Present，到场
如果学生能够 同时 满足下面两个条件，则可以获得出勤奖励：

按 总出勤 计，学生缺勤（'A'）严格 少于两天。
学生 不会 存在 连续 3 天或 连续 3 天以上的迟到（'L'）记录。
给你一个整数 n ，表示出勤记录的长度（次数）。请你返回记录长度为 n 时，可能获得出勤奖励的记录情况 数量 。
答案可能很大，所以返回对 109 + 7 取余 的结果。
*/
func checkRecord(n int) int {

	f := make([][]int, n)

	for i := 0; i < n; i++ {
		f[i] = make([]int, 3)
	}
	f[0][0] = 1
	f[0][1] = 1
	f[0][2] = 0

	// L P  LL
	mod := 1_000_000_007
	for i := 1; i < n-1; i++ {
		f[i][0] = f[i-1][1] % mod
		f[i][2] = (f[i-1][0] + f[i-1][2]) % mod
		if i == 1 {
			f[1][2] = 1
		} else {
			f[i][2] = f[i-1][1] % mod
		}
	}

	s := 0

	if n == 1 {
		return 3
	}
	for i := 0; i < 3; i++ {
		s += f[n-2][i] * n
		s %= mod
	}
	return s
}

func checkRecord2(n int) (ans int) {
	const mod int = 1e9 + 7
	// 三个维度分别表示：长度，A 的数量，结尾连续 L 的数量
	dp := make([][2][3]int, n+1)
	dp[0][0][0] = 1
	for i := 1; i <= n; i++ {
		// 以 P 结尾的数量
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 2; k++ {
				dp[i][j][0] = (dp[i][j][0] + dp[i-1][j][k]) % mod
			}
		}
		// 以 A 结尾的数量
		for k := 0; k <= 2; k++ {
			dp[i][1][0] = (dp[i][1][0] + dp[i-1][0][k]) % mod
		}
		// 以 L 结尾的数量
		for j := 0; j <= 1; j++ {
			for k := 1; k <= 2; k++ {
				dp[i][j][k] = (dp[i][j][k] + dp[i-1][j][k-1]) % mod
			}
		}
	}
	for j := 0; j <= 1; j++ {
		for k := 0; k <= 2; k++ {
			ans = (ans + dp[n][j][k]) % mod
		}
	}
	return ans
}

func main() {
	println(checkRecord(2))
}
