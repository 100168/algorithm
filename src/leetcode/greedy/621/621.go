package main

import (
	"container/heap"
	"fmt"
)

/**
给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表，用字母 A 到 Z 表示，以及一个冷却时间 n。
每个周期或时间间隔允许完成一项任务。任务可以按任何顺序完成，但有一个限制：两个 相同种类 的任务之间必须有长度为 n 的冷却时间。

返回完成所有任务所需要的 最短时间间隔 。


示例 1：

输入：tasks = ["A","A","A","B","B","B"], n = 2
输出：8
解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B
     在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。
*/

func leastInterval(tasks []byte, n int) int {

	cnt := make(map[byte]int)

	for _, v := range tasks {
		cnt[v]++
	}

	h := new(hp)
	for b, v := range cnt {

		heap.Push(h, pair{b, v})

	}

	ans := 0

	for h.Len() > 0 {

		cc := 0
		curPair := make([]pair, 0)
		for i := 0; i <= n && h.Len() > 0; i++ {
			c := heap.Pop(h).(pair)
			c.v--
			if c.v > 0 {
				curPair = append(curPair, c)
			}
			cc++
		}

		if len(curPair) > 0 {
			ans += n + 1
		} else {
			ans += cc
		}

		for _, v := range curPair {
			heap.Push(h, v)
		}
	}
	return ans

	//怎么判断是否合法

}

func leastInterval2(tasks []byte, n int) int {

	cnt := make(map[byte]int)

	mx := 0
	for _, v := range tasks {
		cnt[v]++
		mx = max(mx, cnt[v])
	}

	c := 0
	for _, v := range cnt {

		if v == mx {
			c++
		}
	}
	return max(len(tasks), (mx-1)*(n+1)+c)

}

//class Solution {
//    public int leastInterval(char[] tasks, int n) {
//        int[] temp = new int[26];
//        int countMaxTask = 0;
//        int maxTask=0;
//        for(char c:tasks){
//            temp[c-'A']++;
//            maxTask = Math.max(temp[c-'A'],maxTask);
//        }
//        for(int i=0;i<26;i++){
//            if(temp[i]==maxTask){
//                countMaxTask++;
//            }
//        }
//        return Math.max(tasks.length,(maxTask-1)*(n+1)+countMaxTask);
//    }
//}

type pair struct {
	x byte
	v int
}
type hp []pair

func (h *hp) Less(i, j int) bool {

	return (*h)[i].v > (*h)[j].v || (*h)[i].v == (*h)[j].v && (*h)[i].x < (*h)[j].x
}

func (h *hp) Swap(i, j int) {

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *hp) Push(v any) {
	*h = append(*h, v.(pair))
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
}
