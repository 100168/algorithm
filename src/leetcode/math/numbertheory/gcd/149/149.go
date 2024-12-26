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
	type pair struct {
		x, y int
	}

	for i := 0; i < len(points); i++ {
		cnt := make(map[pair]int)
		for j := i + 1; j < len(points); j++ {
			k1, k2 := points[j][1]-points[i][1], points[j][0]-points[i][0]
			g1 := gcd(k1, k2)
			k1 /= g1
			k2 /= g1
			cnt[pair{k1, k2}]++
			ans = max(ans, cnt[pair{k1, k2}])
		}
	}
	return ans + 1
}

func gcd(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

/*
*
扩展欧几里得

1.求系数
2.求gcd
3.求逆元   a跟b互质    a%b==1   =>= xa+yb=1
4.求丢番图方程解

gcd(a,b) = ax1+by1

gcd(a,a%b) =bx2+(a-(a/b)*b)*y2 = a*y2-b(x2-(a/b)*y2)  =>x1 = y2  y1=x2-(a/b)*y2
*/
func exgcd(a, b int) (int, int) {
	if b == 0 {
		return 1, 0
	} else {
		x1, y1 := exgcd(b, a%b)
		x := y1
		y := x1 - y1*(a/b)
		return x, y
	}
}

const mod = int(1e9 + 7)

func fermat(a, b int) int {

	x := 1

	for ; b > 0; b >>= 1 {

		if b&1 != 0 {
			x = x * a % mod
		}
		a = a * a % mod
	}
	return x
}

func main() {
	//fmt.Println(maxPoints([][]int{{0, 0}, {4, 5}, {7, 8}, {8, 9}, {5, 6}, {3, 4}, {1, 1}}))

	for i := 1; i <= 10000; i++ {
		g1, _ := exgcd(i, mod)
		g1 = (g1 + mod) % mod
		g2 := fermat(i, mod-2)
		if g1 != g2 {
			fmt.Println(i)
		}
	}

}
