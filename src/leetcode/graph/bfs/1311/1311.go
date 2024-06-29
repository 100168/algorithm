package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
*
有 n 个人，每个人都有一个  0 到 n-1 的唯一 id 。

给你数组 watchedVideos  和 friends ，
其中 watchedVideos[i]  和 friends[i] 分别表示 id = i 的人观看过的视频列表和他的好友列表。

Level 1 的视频包含所有你好友观看过的视频，level 2 的视频包含所有你好友的好友观看过的视频，
以此类推。一般的，Level 为 k 的视频包含所有从你出发，最短距离为 k 的好友观看过的视频。

给定你的 id  和一个 level 值，请你找出所有指定 level 的视频，
并将它们按观看频率升序返回。如果有频率相同的视频，请将它们按字母顺序从小到大排列。

输入：watchedVideos = [["A","B"],["C"],["B","C"],["D"]], friends = [[1,2],[0,3],[0,3],[1,2]], id = 0, level = 1
输出：["B","C"]
解释：
你的 id 为 0（绿色），你的朋友包括（黄色）：
id 为 1 -> watchedVideos = ["C"]
id 为 2 -> watchedVideos = ["B","C"]
你朋友观看过视频的频率为：
B -> 1
C -> 2
*/
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {

	n := len(watchedVideos)

	var queue []int

	visited := make([]bool, n)

	dis := make([]int, n)
	queue = append(queue, id)
	visited[id] = true
	deep := 0
	for len(queue) > 0 {
		deep++
		temp := queue
		queue = nil
		for _, v := range temp {
			for _, next := range friends[v] {
				if !visited[next] {
					visited[next] = true
					dis[next] = deep
					queue = append(queue, next)
				}
			}
		}
	}
	cnt := make(map[string]int)
	ans := make([]string, 0)

	for i, v := range dis {
		if v == level {
			for _, w := range watchedVideos[i] {
				cnt[w]++
			}
		}
	}
	for s := range cnt {
		ans = append(ans, s)
	}
	slices.SortStableFunc(ans, func(a, b string) int {
		if cnt[a] == cnt[b] {
			return strings.Compare(a, b)
		}
		return cnt[a] - cnt[b]
	})
	return ans

}

func watchedVideosByFriends2(watchedVideos [][]string, friends [][]int, id int, level int) []string {

	n := len(watchedVideos)

	deeps := make([]int, n)

	for i := range deeps {
		deeps[i] = n
	}
	deeps[id] = 0
	cnt := make(map[string]int)
	var dfs func(int, int)

	dfs = func(x int, deep int) {
		if deeps[x] < deep {
			return
		}
		deeps[x] = deep
		for _, v := range friends[x] {
			dfs(v, deep+1)
		}
	}
	dfs(id, 0)
	ans := make([]string, 0)

	for i, v := range deeps {
		if v == level {
			for _, w := range watchedVideos[i] {
				cnt[w]++
			}
		}
	}
	for s := range cnt {
		ans = append(ans, s)
	}
	slices.SortStableFunc(ans, func(a, b string) int {
		if cnt[a] == cnt[b] {
			return strings.Compare(a, b)
		}
		return cnt[a] - cnt[b]
	})
	return ans
}

func main() {
	fmt.Println(watchedVideosByFriends([][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}, [][]int{{1, 2, 3}, {0, 2, 3}, {0, 1, 3}, {0, 1, 2}}, 0, 1))
}
