package main

import (
	"fmt"
	"math"
)

/*
*
给定一个数组 books ，其中 books[i] = [thicknessi, heighti] 表示第 i 本书的厚度和高度。你也会得到一个整数 shelfWidth 。

按顺序 将这些书摆放到总宽度为 shelfWidth 的书架上。

先选几本书放在书架上（它们的厚度之和小于等于书架的宽度 shelfWidth ），然后再建一层书架。重复这个过程，直到把所有的书都放在书架上。

需要注意的是，在上述过程的每个步骤中，摆放书的顺序与给定图书数组 books 顺序相同。

例如，如果这里有 5 本书，那么可能的一种摆放情况是：第一和第二本书放在第一层书架上，第三本书放在第二层书架上，第四和第五本书放在最后一层书架上。
每一层所摆放的书的最大高度就是这一层书架的层高，书架整体的高度为各层高之和。

以这种方式布置书架，返回书架整体可能的最小高度。
*/
func minHeightShelves(books [][]int, shelfWidth int) int {

	n := len(books)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int

	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}
		cur := math.MaxInt
		sw := 0
		sh := 0
		for j := i; j >= 0; j-- {
			c := books[j]
			cw, ch := c[0], c[1]
			sh = max(sh, ch)
			sw += cw
			if sw > shelfWidth {
				break
			}
			cur = min(cur, dfs(j-1)+sh)
		}
		memo[i] = cur
		return cur

	}
	return dfs(n - 1)
}

func main() {

	fmt.Println(minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4))
}
