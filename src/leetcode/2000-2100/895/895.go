package main

import "container/heap"

/**
设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。

实现 FreqStack 类:

FreqStack() 构造一个空的堆栈。
void push(int val) 将一个整数 val 压入栈顶。
int pop() 删除并返回堆栈中出现频率最高的元素。
如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素。

示例 1：

输入：
["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
输出：[null,null,null,null,null,null,null,5,7,5,4]
解释：
FreqStack = new FreqStack();
freqStack.push (5);//堆栈为 [5]
freqStack.push (7);//堆栈是 [5,7]
freqStack.push (5);//堆栈是 [5,7,5]
freqStack.push (7);//堆栈是 [5,7,5,7]
freqStack.push (4);//堆栈是 [5,7,5,7,4]
freqStack.push (5);//堆栈是 [5,7,5,7,4,5]
freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
freqStack.pop ();//返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
freqStack.pop ();//返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。

思路:
1.堆+时间搓
2.队列套彰
*/

type FreqStack struct {
	cntMap   map[int]int
	indexMap map[int][]int
	count    int
	hp       *myHeap
}

func Constructor() FreqStack {
	fq := new(FreqStack)
	fq.hp = new(myHeap)
	fq.indexMap = make(map[int][]int)
	fq.cntMap = make(map[int]int)
	return *fq
}

func (t *FreqStack) Push(val int) {
	t.count++
	indexes := t.indexMap[val]
	indexes = append(indexes, t.count)
	t.indexMap[val] = indexes
	t.cntMap[val]++
	heap.Push(t.hp, pair{val, t.cntMap[val], indexes[len(indexes)-1]})

}

func (t *FreqStack) Pop() int {

	c := heap.Pop(t.hp).(pair)
	for len(t.indexMap[c.val]) == 0 || c.idx != t.indexMap[c.val][len(t.indexMap[c.val])-1] {
		c = heap.Pop(t.hp).(pair)
	}
	if t.cntMap[c.val]--; t.cntMap[c.val] == 0 {
		delete(t.cntMap, c.val)
		delete(t.indexMap, c.val)
		return c.val
	}

	c.cnt--
	t.indexMap[c.val] = t.indexMap[c.val][:len(t.indexMap[c.val])-1]
	c.idx = t.indexMap[c.val][len(t.indexMap[c.val])-1]
	heap.Push(t.hp, c)
	return c.val
}

type pair struct {
	val int
	cnt int
	idx int
}
type myHeap []pair

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i].cnt > (*h)[j].cnt || (*h)[i].cnt == (*h)[j].cnt && (*h)[i].idx > (*h)[j].idx
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v any) {
	*h = append(*h, v.(pair))
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
