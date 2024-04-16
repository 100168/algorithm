package main

func numberOfWays(s string) int64 {

	m := len(s)

	sum := make([]int, m+1)
	ans := int64(0)

	for i := 1; i <= m; i++ {
		sum[i] = sum[i-1] + int(s[i-1]-'0')
	}

	for i := 1; i < m-1; i++ {
		if s[i] == '0' {
			pre := sum[i]
			suff := sum[m] - pre
			ans += int64(pre * suff)
		} else {
			pre := i - sum[i]
			suff := (m - i - 1) - (sum[m] - sum[i+1])
			ans += int64(pre * suff)
		}
	}
	return ans

}

func numberOfWays2(s string) int64 {
	var dp0, dp1 [3]int
	for _, c := range s {
		if c == '0' {
			dp0[2] += dp1[1]
			dp0[1] += dp1[0]
			dp0[0]++
		} else {
			dp1[2] += dp0[1]
			dp1[1] += dp0[0]
			dp1[0]++
		}
	}
	return int64(dp0[2] + dp1[2])
}
