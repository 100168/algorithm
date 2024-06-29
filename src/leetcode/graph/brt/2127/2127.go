package main

import "fmt"

/*
*一个公司准备组织一场会议，邀请名单上有 n 位员工。公司准备了一张 圆形 的桌子，可以坐下 任意数目 的员工。

员工编号为 0 到 n - 1 。每位员工都有一位 喜欢 的员工，每位员工 当且仅当 他被安排在喜欢员工的旁边，
他才会参加会议。每位员工喜欢的员工 不会 是他自己。

给你一个下标从 0 开始的整数数组 favorite ，其中 favorite[i] 表示第 i 位员工喜欢的员工。
请你返回参加会议的 最多员工数目 。
思路：

1.基环树

分两种情况：
1.基环树长度大于2最大值则为基环树长度
2.基环树长度为2，把所有长度相加，连在一起
3.取两者最大值
*/
func maximumInvitations(favorite []int) int {

	n := len(favorite)

	inDegree := make([]int, n)
	for _, v := range favorite {
		inDegree[v]++
	}

	var queue []int
	for i := range inDegree {

		if inDegree[i] == 0 {

			queue = append(queue, i)
		}
	}

	//先走一遍拓扑排序，并且记录每个点的深度
	deep := make([]int, n)
	for len(queue) > 0 {
		x := queue[0]
		deep[x]++
		queue = queue[1:]
		y := favorite[x]
		inDegree[y]--
		deep[y] = deep[x]
		if inDegree[y] == 0 {
			queue = append(queue, y)
		}
	}

	ans := 0

	sunRange := 0
	for i := range inDegree {
		if inDegree[i] == 0 {
			continue
		}

		cnt := 1
		for x := favorite[i]; x != i; x = favorite[x] {
			cnt++
			inDegree[x]--
		}
		if cnt == 2 {
			sunRange += 2 + deep[i] + deep[favorite[i]]
		} else {
			ans = max(ans, cnt)
		}

	}
	return max(sunRange, ans)

}

func main() {
	fmt.Println(maximumInvitations([]int{2, 2, 1, 2}))
}
