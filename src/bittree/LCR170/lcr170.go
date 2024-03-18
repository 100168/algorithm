package main

import (
	"slices"
	"sort"
)

// 在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的
// 「交易逆序对」总数。
//
// 示例 1:
//
// 输入：record = [9, 7, 5, 4, 6]
// 输出：8
// 解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。
//
// 限制：
//
// 0 <= record.length <= 50000
//
// Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 1093 👎 0
type bitSet struct {
	sum []int
	len int
}

func lowBit(index int) int {
	return index & -index
}

func (b *bitSet) query(index int) int {
	ans := 0
	for index > 0 {
		ans += b.sum[index]
		index -= lowBit(index)
	}
	return ans
}
func (b *bitSet) update(index, val int) {

	for index < b.len {
		b.sum[index] += val
		index += lowBit(index)
	}
}

func reversePairs(record []int) int {

	type pair struct {
		index int
		val   int
	}
	n := len(record)
	sortR := make([]pair, n)
	for i := range record {
		sortR[i] = pair{i, record[i]}
	}
	sort.Slice(sortR, func(i, j int) bool {
		return sortR[i].val < sortR[j].val
	})

	b := new(bitSet)
	//表示排序好的下标
	sa := make([]int, n)
	for i := range sa {
		sa[i] = i
	}
	slices.SortStableFunc(sa, func(i, j int) int { return record[i] - record[j] })

	b.sum = make([]int, n+1)
	b.len = len(record) + 1
	ans := 0
	for i := 0; i < len(record); i++ {
		v := sa[i] + 1
		ans += b.query(n) - b.query(v)
		b.update(v, 1)
	}
	return ans

}

func main() {
	println(reversePairs([]int{7, 5, 6, 4}))
}
