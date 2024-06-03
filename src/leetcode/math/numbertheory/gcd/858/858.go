package main

import "fmt"

/*
有一个特殊的正方形房间，每面墙上都有一面镜子。除西南角以外，每个角落都放有一个接受器，编号为 0， 1，以及 2。

正方形房间的墙壁长度为 p，一束激光从西南角射出，首先会与东墙相遇，入射点到接收器 0 的距离为 q 。

返回光线最先遇到的接收器的编号（保证光线最终会遇到一个接收器）。
*/
func mirrorReflection(p int, q int) int {

	if q == 0 {
		return 0
	}
	lcm := p * q / gcd(p, q)

	mulQ := lcm / q
	mulP := lcm / p

	if mulQ%2 == 0 {
		return 2
	} else {

		if mulP%2 == 0 {
			return 0
		}
		return 1
	}

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	//fmt.Println(mirrorReflection(2, 1))
	fmt.Println(mirrorReflection(3, 2))
}
