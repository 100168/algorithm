package main

/**

在一个无限的 x 坐标轴上，有许多水果分布在其中某些位置。给你一个二维整数数组 fruits ，
其中 fruits[i] = [positioni, amounti] 表示共有 amounti 个水果放置在 positioni 上。
fruits 已经按 positioni 升序排列 ，每个 positioni 互不相同 。
另给你两个整数 startPos 和 k 。最初，你位于 startPos 。从任何位置，
你可以选择 向左或者向右 走。在 x 轴上每移动 一个单位 ，就记作 一步 。
你总共可以走 最多 k 步。你每达到一个位置，都会摘掉全部的水果，水果也将从该位置消失（不会再生）。

返回你可以摘到水果的 最大总数 。*/

func maxTotalFruits(fruits [][]int, startPos int, k int) int {

	fruitMap := make(map[int]int)

	for _, v := range fruits {
		fruitMap[v[0]] = v[1]
	}
	w := 0
	ans := 0

	l := max(0, startPos-k)

	for r := max(0, startPos-k); r <= startPos+k; r++ {
		w += fruitMap[r]
		for min((startPos-l)*2+(r-startPos), (r-startPos)*2+startPos-l) > k {
			w -= fruitMap[l]
			l++
		}

		ans = max(w, ans)
	}
	return ans
}

func main() {

}
