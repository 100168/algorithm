package main

/*给你一个 n 个视频的上传序列，每个视频编号为 1 到 n 之间的 不同 数字，你需要依次将这些视频上传到服务器。请你实现一个数据结构，在上传的过程中计算 最长上传前缀 。

如果 闭区间 1 到 i 之间的视频全部都已经被上传到服务器，那么我们称 i 是上传前缀。最长上传前缀指的是符合定义的 i 中的 最大值 。

请你实现 LUPrefix 类：

LUPrefix(int n) 初始化一个 n 个视频的流对象。
void upload(int video) 上传 video 到服务器。
int longest() 返回上述定义的 最长上传前缀 的长度。*/

type LUPrefix struct {
	unitFind *UnitFind
}

func Constructor(n int) LUPrefix {

	lup := new(LUPrefix)
	lup.unitFind = new(UnitFind)
	lup.unitFind.init(n + 2)
	return *lup

}

func (this *LUPrefix) Upload(video int) {

	unitFind := this.unitFind
	unitFind.find(video)
	if unitFind.parent[video-1] != 0 {
		unitFind.unit(video, video-1)
	}
	if unitFind.parent[video+1] != 0 {
		unitFind.unit(video, video+1)
	}

}

func (this *LUPrefix) Longest() int {
	return this.unitFind.maxLen
}

type UnitFind struct {
	size   []int
	parent []int
	maxLen int
}

func (unitFind *UnitFind) init(n int) {
	unitFind.parent = make([]int, n+1)
	unitFind.size = make([]int, n+1)
}

func (unitFind *UnitFind) find(cur int) int {

	if unitFind.parent[cur] == 0 {
		unitFind.parent[cur] = 1
		unitFind.size[cur] = 1
		if cur == 1 {
			unitFind.maxLen = 1
		}
		return cur
	}

	for cur != unitFind.parent[cur] {
		unitFind.parent[cur] = unitFind.parent[unitFind.parent[cur]]
		cur = unitFind.parent[cur]
	}
	return cur

}

func (unitFind *UnitFind) unit(a, b int) {
	findA := unitFind.find(a)
	findB := unitFind.find(b)
	if findA == findB {
		return
	}
	maxFind := findA
	minFind := findA
	if findB > findA {
		maxFind = findB
	} else {
		minFind = findB
	}
	unitFind.size[maxFind] += unitFind.size[minFind]
	unitFind.size[minFind] = 0
	unitFind.parent[minFind] = maxFind
	if unitFind.size[maxFind] == maxFind {
		unitFind.maxLen = maxFind
	}
}

type Bit struct {
	sum    []int
	n      int
	maxLen int
}

func (bit *Bit) init(n int) {
	bit.sum = make([]int, n+1)
	bit.n = n + 1
}

func lowBit(index int) int {
	return index & -index
}

func (bit *Bit) update(index int) {
	for ; index < bit.n; index += lowBit(index) {
		bit.sum[index]++
	}
}

func main() {

}

/**
 * Your LUPrefix object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.Upload(video);
 * param_2 := obj.Longest();
 */
