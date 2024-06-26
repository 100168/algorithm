package main

import "fmt"

/*
*给你一个有 n 个服务器的计算机网络，服务器编号为 0 到 n - 1 。同时给你一个二维整数数组 edges ，
其中 edges[i] = [ui, vi] 表示服务器 ui 和 vi 之间有一条信息线路，
在 一秒 内它们之间可以传输 任意 数目的信息。再给你一个长度为 n 且下标从 0 开始的整数数组 patience 。

题目保证所有服务器都是 相通 的，也就是说一个信息从任意服务器出发，都可以通过这些信息线路直接或间接地到达任何其他服务器。

编号为 0 的服务器是 主 服务器，其他服务器为 数据 服务器。每个数据服务器都要向主服务器发送信息，
并等待回复。信息在服务器之间按 最优 线路传输，也就是说每个信息都会以 最少时间 到达主服务器。
主服务器会处理 所有 新到达的信息并 立即 按照每条信息来时的路线 反方向 发送回复信息。

在 0 秒的开始，所有数据服务器都会发送各自需要处理的信息。从第 1 秒开始，
每 一秒最 开始 时，每个数据服务器都会检查它是否收到了主服务器的回复信息（包括新发出信息的回复信息）：

如果还没收到任何回复信息，那么该服务器会周期性 重发 信息。
数据服务器 i 每 patience[i] 秒都会重发一条信息，也就是说，
数据服务器 i 在上一次发送信息给主服务器后的 patience[i] 秒 后 会重发一条信息给主服务器。
否则，该数据服务器 不会重发 信息。
当没有任何信息在线路上传输或者到达某服务器时，该计算机网络变为 空闲 状态。

请返回计算机网络变为 空闲 状态的 最早秒数 。
*/
func networkBecomesIdle(edges [][]int, patience []int) int {

	n := len(patience)

	g := make([][]int, n)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	visited := make([]bool, n)
	dis := make([]int, n)

	queue := []int{0}
	visited[0] = true

	deep := 0
	for len(queue) > 0 {
		deep++
		temp := queue
		queue = nil
		for _, x := range temp {
			for _, y := range g[x] {
				if visited[y] {
					continue
				}
				visited[y] = true
				dis[y] = deep
				queue = append(queue, y)
			}

		}
	}

	ans := 0

	//6    4
	for i := 1; i < n; i++ {
		d := dis[i]
		x := patience[i]
		cur := d + d
		mul := (cur - 1) / x

		ans = max(ans, x*mul+cur)
	}
	return ans + 1

}

func main() {
	fmt.Println(networkBecomesIdle([][]int{{0, 1}, {1, 2}}, []int{0, 2, 1}))
}
