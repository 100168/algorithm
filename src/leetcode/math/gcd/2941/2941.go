package main

import "fmt"

/*
*
给定一个整数数组 nums 和一个整数 k.

数组 a 的 gcd-sum 计算方法如下：

设 s 为 a 的所有元素的和。
设 g 为 a 的所有元素的 最大公约数。
a 的 gcd-sum 等于 s * g.
返回 nums 的至少包含 k 个元素的子数组的 最大 gcd-sum。
*/
func maxGcdSum(nums []int, k int) int64 {

	n := len(nums)
	sum := make([]int, n+1)
	type pair struct {
		x, y int
	}
	ans := 0
	var gcdList []pair
	for i, v := range nums {
		sum[i+1] = sum[i] + v
		var newGcdList []pair
		gcdList = append(gcdList, pair{v, i})

		for _, x := range gcdList {
			cur := gcd(v, x.x)
			if i-x.y+1 >= k {
				ans = max(ans, (sum[i+1]-sum[x.y])*cur)
			}
			if len(newGcdList) > 0 && newGcdList[len(newGcdList)-1].x == x.x {
				continue
			}
			newGcdList = append(newGcdList, pair{cur, x.y})
		}
		gcdList = newGcdList

	}
	return int64(ans)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(maxGcdSum([]int{2, 1, 4, 4, 4, 2}, 2))
}
