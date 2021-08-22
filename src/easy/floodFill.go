package main

import "math"

/*
有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。

给你一个坐标(sr, sc)表示图像渲染开始的像素值（行 ，列）和一个新的颜色值newColor，让你重新上色这幅图像。

为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。

最后返回经过上色渲染后的图像。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flood-fill
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	n := len(image)
	m := len(image[0])

	infect(image, sr, sc, n-1, m-1, image[sr][sc])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if image[i][j] == math.MaxInt64 {
				image[i][j] = newColor
			}
		}
	}
	return image
}
func infect(image [][]int, sr int, sc int, n int, m int, currentColor int) {

	if image[sr][sc] == currentColor {
		image[sr][sc] = math.MaxInt64
	} else {
		return
	}
	if sr+1 <= n {
		infect(image, sr+1, sc, n, m, currentColor)
	}
	if sr-1 >= 0 {
		infect(image, sr-1, sc, n, m, currentColor)
	}

	if sc+1 <= m {
		infect(image, sr, sc+1, n, m, currentColor)
	}
	if sc-1 >= 0 {
		infect(image, sr, sc-1, n, m, currentColor)
	}
}
func main() {

}
