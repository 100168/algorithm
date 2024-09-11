package main

import "fmt"

/**
给你一个下标从 0 开始、大小为 m x n 的网格 image ，表示一个灰度图像，
其中 image[i][j] 表示在范围 [0..255] 内的某个像素强度。另给你一个 非负 整数 threshold 。

如果 image[a][b] 和 image[c][d] 满足 |a - c| + |b - d| == 1 ，则称这两个像素是 相邻像素 。

区域 是一个 3 x 3 的子网格，且满足区域中任意两个 相邻 像素之间，像素强度的 绝对差 小于或等于 threshold 。

区域 内的所有像素都认为属于该区域，而一个像素 可以 属于 多个 区域。

你需要计算一个下标从 0 开始、大小为 m x n 的网格 result ，其中 result[i][j] 是 image[i][j] 所属区域的 平均 强度，
向下取整 到最接近的整数。如果 image[i][j] 属于多个区域，result[i][j] 是这些区域的 “取整后的平均强度” 的 平均值，
也 向下取整 到最接近的整数。如果 image[i][j] 不属于任何区域，则 result[i][j] 等于 image[i][j] 。

返回网格 result 。
*/

func resultGrid(image [][]int, threshold int) [][]int {

	m := len(image)
	n := len(image[0])

	res := make([][]int, m)

	cnt := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		cnt[i] = make([]int, n)
	}

	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

	for i := 2; i < m; i++ {
	next:
		for j := 2; j < n; j++ {
			t := 0
			for x := i - 2; x <= i; x++ {
				for y := j - 2; y <= j; y++ {
					for _, d := range dir {
						nx, ny := x+d[0], y+d[1]
						if nx >= i-2 && nx <= i && ny >= j-2 && ny <= j {
							if abs(image[nx][ny]-image[x][y]) > threshold {
								continue next
							}
						}
					}
					t += image[x][y]
				}
			}

			for x := i - 2; x <= i; x++ {
				for y := j - 2; y <= j; y++ {
					cnt[x][y]++
					res[x][y] += t / 9
				}
			}

		}
	}

	for i, r := range res {
		for j, v := range r {
			if cnt[i][j] > 0 {
				res[i][j] = v / cnt[i][j]
			} else {
				res[i][j] = image[i][j]
			}
		}
	}
	return res

}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(resultGrid([][]int{{5, 6, 7, 10}, {8, 9, 10, 10}, {11, 12, 13, 10}}, 3))
}
