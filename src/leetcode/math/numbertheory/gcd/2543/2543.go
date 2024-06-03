package main

/*
*
给你一个无穷大的网格图。一开始你在 (1, 1) ，你需要通过有限步移动到达点 (targetX, targetY) 。

每一步 ，你可以从点 (x, y) 移动到以下点之一：
(x, y - x)
(x - y, y)
(x+x, y)
(x, y + y)
给你两个整数 targetX 和 targetY ，分别表示你最后需要到达点的 X 和 Y 坐标。如果你可以从 (1, 1) 出发到达这个点，请你返回true ，否则返回 false 。

思路：
1.先将targetX和targetY 变成奇数
2.然后再根据裴蜀定理如果两个数互质则肯定存在a,b使得x-y=1或y-x=1
*/

func isReachable(targetX int, targetY int) bool {
	return gcd(targetX, targetY)&(gcd(targetX, targetY)-1) == 0
}

func isReachable2(targetX int, targetY int) bool {

	for targetX%2 == 0 {
		targetX = targetX / 2
	}
	for targetY%2 == 0 {
		targetY = targetY / 2
	}
	return gcd(targetX, targetY) == 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
