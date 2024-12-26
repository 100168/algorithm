package main

import "fmt"

/*
*
现有一个下标从 1 开始的 8 x 8 棋盘，上面有 3 枚棋子。

给你 6 个整数 a 、b 、c 、d 、e 和 f ，其中：

(a, b) 表示白色车的位置。
(c, d) 表示白色象的位置。
(e, f) 表示黑皇后的位置。
假定你只能移动白色棋子，返回捕获黑皇后所需的最少移动次数。

请注意：

车可以向垂直或水平方向移动任意数量的格子，但不能跳过其他棋子。
象可以沿对角线方向移动任意数量的格子，但不能跳过其他棋子。
如果车或象能移向皇后所在的格子，则认为它们可以捕获皇后。
皇后不能移动。
*/
func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {

	if a == e {
		if c == a && d > min(b, f) && d < max(b, f) {
			return 2
		} else {
			return 1
		}
	}

	if b == f {
		if d == b && c > min(a, e) && c < max(a, e) {
			return 2
		} else {
			return 1
		}
	}

	if abs(f-d) == abs(e-c) {
		if (f-d)/(e-c) == (f-b)/(e-a) && (a > min(c, e) && a < max(c, e)) {
			return 2
		}
		return 1

	}

	return 2

}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(minMovesToCaptureTheQueen(8, 4, 8, 8, 7, 7))
}
