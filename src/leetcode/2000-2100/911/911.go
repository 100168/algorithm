package main

import "fmt"

/**
给你两个整数数组 persons 和 times 。在选举中，第 i 张票是在时刻为 times[i] 时投给候选人 persons[i] 的。

对于发生在时刻 t 的每个查询，需要找出在 t 时刻在选举中领先的候选人的编号。

在 t 时刻投出的选票也将被计入我们的查询之中。在平局的情况下，最近获得投票的候选人将会获胜。

实现 TopVotedCandidate 类：

TopVotedCandidate(int[] persons, int[] times) 使用 persons 和 times 数组初始化对象。
int q(int t) 根据前面描述的规则，返回在时刻 t 在选举中领先的候选人的编号。

思路：
1.预处理
2.二分
*/

type TopVotedCandidate struct {
	t  []int
	rk map[int]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	tp := new(TopVotedCandidate)
	tp.t = times
	tp.rk = make(map[int]int)

	cntMap := make(map[int]int)
	curMax := -1
	for i, v := range times {
		cntMap[persons[i]]++
		if cntMap[persons[i]] >= cntMap[curMax] || curMax == -1 {
			curMax = persons[i]
		}
		tp.rk[v] = curMax
	}
	return *tp
}

//[[[0,1,1,0,0,1,0],[0,5,10,15,20,25,30]],[3],[12],[25],[15],[24],[8]]

func (tp *TopVotedCandidate) Q(t int) int {
	tt := tp.t
	rk := tp.rk
	l, r := 0, len(tt)-1
	for l <= r {
		m := (r + l) / 2
		if tt[m] <= t {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return rk[tt[r]]
}

func main() {

	topVotedCandidate := Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})
	fmt.Println(topVotedCandidate.Q(12))

}
