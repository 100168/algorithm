package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个 正整数 数组 beans ，其中每个整数表示一个袋子里装的魔法豆的数目。

请你从每个袋子中 拿出 一些豆子（也可以 不拿出），使得剩下的 非空 袋子中（即 至少还有一颗 魔法豆的袋子）魔法豆的数目 相等。
一旦把魔法豆从袋子中取出，你不能再将它放到任何袋子中。

请返回你需要拿出魔法豆的 最少数目。

输入：beans = [4,1,6,5]
输出：4
解释：
  - 我们从有 1 个魔法豆的袋子中拿出 1 颗魔法豆。
    剩下袋子中魔法豆的数目为：[4,0,6,5]
  - 然后我们从有 6 个魔法豆的袋子中拿出 2 个魔法豆。
    剩下袋子中魔法豆的数目为：[4,0,4,5]
  - 然后我们从有 5 个魔法豆的袋子中拿出 1 个魔法豆。
    剩下袋子中魔法豆的数目为：[4,0,4,4]

总共拿出了 1 + 2 + 1 = 4 个魔法豆，剩下非空袋子中魔法豆的数目相等。
没有比取出 4 个魔法豆更少的方案。

思路；先排序再枚举
*/
func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)

	n := len(beans)
	sum := make([]int, n+1)
	for i, v := range beans {
		sum[i+1] = sum[i] + v
	}

	ans := math.MaxInt
	for i, v := range beans {
		ans = min(ans, sum[i]+sum[n]-sum[i]-(n-i)*v)
	}

	return int64(ans)

}

func main() {
	fmt.Println(minimumRemoval([]int{63, 43, 12, 94}))
}
