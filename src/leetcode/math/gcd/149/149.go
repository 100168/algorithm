package main

import "fmt"

/*
*
给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。

points = [[1,1],[2,2],[3,3]]
*/
func maxPoints(points [][]int) int {

	ans := 0

	if len(points) == 1 {
		return 1
	}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			cnt := 2
			k1, k2 := points[j][1]-points[i][1], points[j][0]-points[i][0]
			g1 := gcd(k1, k2)
			k1 /= g1
			k2 /= g1
			for k := j + 1; k < len(points); k++ {
				k11, k22 := points[k][1]-points[i][1], points[k][0]-points[i][0]
				g2 := gcd(k11, k22)
				k11 /= g2
				k22 /= g2
				if k11 == k1 && k22 == k2 {
					cnt++
				}
			}
			ans = max(ans, cnt)
		}
	}
	return ans
}

func gcd(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(maxPoints([][]int{{0, 0}, {4, 5}, {7, 8}, {8, 9}, {5, 6}, {3, 4}, {1, 1}}))
}
