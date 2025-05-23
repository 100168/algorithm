package main

import (
	"container/heap"
	"fmt"
)

/*
给你一个字符串 initialCurrency，表示初始货币类型，并且你一开始拥有 1.0 单位的 initialCurrency。

另给你四个数组，分别表示货币对（字符串）和汇率（实数）：

pairs1[i] = [startCurrencyi, targetCurrencyi] 表示在 第 1 天，可以按照汇率 rates1[i] 将 startCurrencyi 转换为 targetCurrencyi。
pairs2[i] = [startCurrencyi, targetCurrencyi] 表示在 第 2 天，可以按照汇率 rates2[i] 将 startCurrencyi 转换为 targetCurrencyi。
此外，每种 targetCurrency 都可以以汇率 1 / rate 转换回对应的 startCurrency。
你可以在 第 1 天 使用 rates1 进行任意次数的兑换（包括 0 次），然后在 第 2 天 使用 rates2 再进行任意次数的兑换（包括 0 次）。

返回在两天兑换后，最大可能拥有的 initialCurrency 的数量。

注意：汇率是有效的，并且第 1 天和第 2 天的汇率之间相互独立，不会产生矛盾。



示例 1：

输入： initialCurrency = "EUR", pairs1 = [["EUR","USD"],["USD","JPY"]], rates1 = [2.0,3.0], pairs2 = [["JPY","USD"],["USD","CHF"],["CHF","EUR"]], rates2 = [4.0,5.0,6.0]

输出： 720.00000

解释：

根据题目要求，需要最大化最终的 EUR 数量，从 1.0 EUR 开始：

第 1 天：
将 EUR 换成 USD，得到 2.0 USD。
将 USD 换成 JPY，得到 6.0 JPY。
第 2 天：
将 JPY 换成 USD，得到 24.0 USD。
将 USD 换成 CHF，得到 120.0 CHF。
最后将 CHF 换回 EUR，得到 720.0 EUR。
示例 2：

输入： initialCurrency = "NGN", pairs1 = [["NGN","EUR"]], rates1 = [9.0], pairs2 = [["NGN","EUR"]], rates2 = [6.0]

输出： 1.50000

解释：

在第 1 天将 NGN 换成 EUR，并在第 2 天用反向汇率将 EUR 换回 NGN，可以最大化最终的 NGN 数量。

示例 3：

输入： initialCurrency = "USD", pairs1 = [["USD","EUR"]], rates1 = [1.0], pairs2 = [["EUR","JPY"]], rates2 = [10.0]

输出： 1.00000

解释：

在这个例子中，不需要在任何一天进行任何兑换。
*/

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {

	g := make(map[string][]pair)

	for i, v := range pairs1 {

		x, y, v := v[0], v[1], rates1[i]

		g[x] = append(g[x], pair{y, v})
		g[y] = append(g[y], pair{x, 1 / v})

	}
	g2 := make(map[string][]pair)
	for i, v := range pairs2 {

		x, y, v := v[0], v[1], rates2[i]

		g2[x] = append(g2[x], pair{y, v})
		g2[y] = append(g2[y], pair{x, 1 / v})

	}

	dis := make(map[string]float64)

	dis[initialCurrency] = float64(1)

	h := &hp{pair{initialCurrency, float64(1)}}

	for h.Len() > 0 {
		c := heap.Pop(h).(pair)
		if dis[c.x] > c.v {
			continue
		}
		for _, v := range g[c.x] {
			y, d := v.x, c.v*v.v
			if d > dis[y] {
				dis[y] = d
				heap.Push(h, pair{y, d})
			}
		}
	}

	ans := float64(0)
	for k, v := range dis {

		dis2 := make(map[string]float64)
		dis2[k] = v

		h2 := &hp{pair{k, v}}

		for h2.Len() > 0 {
			c := heap.Pop(h2).(pair)
			if dis2[c.x] > c.v {
				continue
			}
			for _, v := range g2[c.x] {
				y, d := v.x, c.v*v.v
				if d > dis2[y] {
					dis2[y] = d
					heap.Push(h2, pair{y, d})
				}
			}
		}

		ans = max(ans, dis2[initialCurrency])

	}
	return ans
}

type pair struct {
	x string

	v float64
}

type hp []pair

func (h *hp) Less(i, j int) bool {

	return (*h)[i].v > (*h)[j].v
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}

func (h *hp) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func main() {
	//fmt.Println(maxAmount("NGN", [][]string{{"NGN", "EUR"}}, []float64{9.0}, [][]string{{"NGN", "EUR"}}, []float64{6.0}))
	fmt.Println(maxAmount("EUR", [][]string{{"EUR", "USD"}, {"USD", "JPY"}}, []float64{2.0, 3.0}, [][]string{{"JPY", "USD"}, {"USD", "CHF"}, {"CHF", "EUR"}}, []float64{4.0, 5.0, 6.0}))
}
