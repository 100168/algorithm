package main

/*
*
给你一个大小为 m x n 、下标从 0 开始的二维矩阵 mat 。在每个单元格，你可以按以下方式生成数字：

最多有 8 条路径可以选择：东，东南，南，西南，西，西北，北，东北。
选择其中一条路径，沿着这个方向移动，并且将路径上的数字添加到正在形成的数字后面。
注意，每一步都会生成数字，例如，如果路径上的数字是 1, 9, 1，那么在这个方向上会生成三个数字：1, 19, 191 。
返回在遍历矩阵所创建的所有数字中，出现频率最高的、大于 10的
质数
；如果不存在这样的质数，则返回 -1 。如果存在多个出现频率最高的质数，那么返回其中最大的那个。
*/
func mostFrequentPrime(mat [][]int) int {

	m, n := len(mat), len(mat[0])

	cntMap := make(map[int]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := mat[i][j]
			if cur > 10 {
				cntMap[cur]++
			}
			for east, next := 1, cur; east+j < n; east++ {
				next = next*10 + mat[i][east+j]
				cntMap[next]++
			}
			for west, next := 1, cur; j-west >= 0; west++ {
				next = next*10 + mat[i][j-west]
				cntMap[next]++
			}
			for north, next := 1, cur; i+north < m; north++ {
				next = next*10 + mat[i+north][j]
				cntMap[next]++
			}
			for south, next := 1, cur; i-south >= 0; south++ {
				next = next*10 + mat[i-south][j]
				cntMap[next]++
			}

			for es, next := 1, cur; i+es < m && j+es < n; es++ {
				next = next*10 + mat[i+es][j+es]
				cntMap[next]++
			}

			for wn, next := 1, cur; i-wn >= 0 && j-wn >= 0; wn++ {
				next = next*10 + mat[i-wn][j-wn]
				cntMap[next]++
			}

			for ws, next := 1, cur; i+ws < m && j-ws >= 0; ws++ {
				next = next*10 + mat[i+ws][j-ws]
				cntMap[next]++
			}
			for en, next := 1, cur; i-en >= 0 && j+en < n; en++ {
				next = next*10 + mat[i-en][j+en]
				cntMap[next]++
			}
		}
	}
	maxFrequent := 0
	ans := -1

	for k, v := range cntMap {
		if isPrime(k) {
			if maxFrequent == v {
				ans = max(ans, k)
			}
			if v > maxFrequent {
				maxFrequent = v
				ans = k
			}

		}
	}
	if ans < 10 {
		return -1
	}
	return ans
}
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func mostFrequentPrime2(mat [][]int) int {
	//优秀的枚举方式
	dirs := []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	m, n := len(mat), len(mat[0])
	cnt := map[int]int{}
	for i, row := range mat {
		for j, v := range row {
			for _, d := range dirs {
				x, y, v := i+d.x, j+d.y, v
				for 0 <= x && x < m && 0 <= y && y < n {
					v = v*10 + mat[x][y]
					// 如果 v 在 cnt 中，那么 v 一定是质数
					if cnt[v] > 0 || isPrime(v) {
						cnt[v]++
					}
					x += d.x
					y += d.y
				}
			}
		}
	}

	ans, maxCnt := -1, 0
	for v, c := range cnt {
		if c > maxCnt {
			ans, maxCnt = v, c
		} else if c == maxCnt {
			ans = max(ans, v)
		}
	}
	return ans
}

func main() {
	println(mostFrequentPrime([][]int{
		{4, 7, 7}}))
}
