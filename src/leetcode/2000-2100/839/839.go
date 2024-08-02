package main

/*
*
如果交换字符串 X 中的两个不同位置的字母，使得它和字符串 Y 相等，那么称 X 和 Y 两个字符串相似。如果这两个字符串本身是相等的，那它们也是相似的。

例如，"tars" 和 "rats" 是相似的 (交换 0 与 2 的位置)； "rats" 和 "arts" 也是相似的，但是 "star" 不与 "tars"，"rats"，或 "arts" 相似。

总之，它们通过相似性形成了两个关联组：{"tars", "rats", "arts"} 和 {"star"}。注意，"tars" 和 "arts" 是在同一组中，即使它们并不相似。形式上，对每个组而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。

给你一个字符串列表 strs。列表中的每个字符串都是 strs 中其它所有字符串的一个字母异位词。请问 strs 中有多少个相似字符串组？

示例 1：

输入：strs = ["tars","rats","arts","star"]
输出：2
示例 2：

输入：strs = ["omv","ovm"]
输出：1
*/
func numSimilarGroups(strs []string) int {

	uf := new(uniFind)
	uf.init(len(strs))

	for i, x := range strs {

		for j, y := range strs {
			if uf.isSameSet(i, j) {
				continue
			}
			diff := 0
			for t := range x {
				if x[t] != y[t] {
					diff++
				}
				if diff > 2 {
					break
				}
			}
			if diff == 2 || diff == 0 {
				uf.unit(i, j)
			}
		}
	}

	ans := 0
	for _, v := range uf.size {
		if v > 0 {
			ans++
		}
	}
	return ans
}

type uniFind struct {
	parent []int
	size   []int
}

func (uniFind *uniFind) init(n int) {
	uniFind.parent = make([]int, n)
	uniFind.size = make([]int, n)
	for i := range uniFind.parent {
		uniFind.parent[i] = i
		uniFind.size[i] = 1
	}
}

func (uniFind *uniFind) find(v int) int {
	parent := uniFind.parent
	for parent[v] != v {
		parent[v] = parent[parent[v]]
		v = parent[v]

	}
	return v
}

func (uniFind *uniFind) isSameSet(a, b int) bool {

	findA := uniFind.find(a)
	findB := uniFind.find(b)
	return findA == findB
}

func (uniFind *uniFind) unit(a, b int) {
	findA := uniFind.find(a)
	findB := uniFind.find(b)
	if findA == findB {
		return
	}
	uniFind.parent[findA] = findB
	uniFind.size[findA] = 0
}
