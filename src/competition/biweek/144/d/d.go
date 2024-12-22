package main

func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)

	// 定义一个三维 dp 数组，初始值为 -1
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	// 初始化起始状态
	dp[0][0][0] = fruits[0][0]

	// 动态规划状态转移
	for step := 1; step < n; step++ {
		for x1 := 0; x1 < n; x1++ {
			for x2 := 0; x2 < n; x2++ {
				for x3 := 0; x3 < n; x3++ {
					y1, y2, y3 := step-x1, step-x2, n-1-step+x3
					if y1 < 0 || y1 >= n || y2 < 0 || y2 >= n || y3 < 0 || y3 >= n {
						continue
					}

					// 当前状态的水果数
					cur := fruits[x1][y1]
					if x1 != x2 || y1 != y2 {
						cur += fruits[x2][y2]
					}
					if (x1 != x3 || y1 != y3) && (x2 != x3 || y2 != y3) {
						cur += fruits[x3][y3]
					}

					// 更新状态
					prevMax := -1
					for dx1 := -1; dx1 <= 0; dx1++ {
						for dx2 := -1; dx2 <= 1; dx2++ {
							for dx3 := -1; dx3 <= 1; dx3++ {
								px1, px2, px3 := x1+dx1, x2+dx2, x3+dx3
								if px1 >= 0 && px1 < n && px2 >= 0 && px2 < n && px3 >= 0 && px3 < n {
									prevMax = max(prevMax, dp[px1][px2][px3])
								}
							}
						}
					}

					if prevMax != -1 {
						dp[x1][x2][x3] = max(dp[x1][x2][x3], prevMax+cur)
					}
				}
			}
		}
	}

	// 返回终点的最大值
	return dp[n-1][n-1][n-1]
}
