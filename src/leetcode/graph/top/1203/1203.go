package main

import "fmt"

/*
*
有 n 个项目，每个项目或者不属于任何小组，或者属于 m 个小组之一。
group[i] 表示第 i 个项目所属的小组，如果第 i 个项目不属于任何小组，则 group[i] 等于 -1。
项目和小组都是从零开始编号的。可能存在小组不负责任何项目，即没有任何项目属于这个小组。

请你帮忙按要求安排这些项目的进度，并返回排序后的项目列表：

同一小组的项目，排序后在列表中彼此相邻。
项目之间存在一定的依赖关系，我们用一个列表 beforeItems 来表示，
其中 beforeItems[i] 表示在进行第 i 个项目前（位于第 i 个项目左侧）应该完成的所有项目。
如果存在多个解决方案，只需要返回其中任意一个即可。如果没有合适的解决方案，就请返回一个 空列表 。
*/
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {

	groupItem := make([][]int, m)
	g := make([][]int, n)
	inDegree := make([]int, n)
	for i := range beforeItems {
		for _, v := range beforeItems[i] {
			g[v] = append(g[v], i)
			inDegree[i]++
		}
	}
	for i, v := range group {
		if v >= 0 {
			groupItem[v] = append(groupItem[v], i)
		}
	}
	ans := make([]int, 0)

	var queue []int

	for i := range inDegree {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		ans = append(ans, cur)

		for _, v := range g[cur] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(ans) != n {
		return nil
	}
	return ans
}

func main() {
	fmt.Println(sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}))
}
