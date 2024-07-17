package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
*
给你一个整数数组 target 。一开始，你有一个数组 A ，它的所有元素均为 1 ，你可以执行以下操作：

令 x 为你数组里所有元素的和
选择满足 0 <= i < target.size 的任意下标 i ，并让 A 数组里下标为 i 处的值为 x 。
你可以重复该过程任意次
如果能从 A 开始构造出目标数组 target ，请你返回 True ，否则返回 False 。

输入：target = [9,3,5]
输出：true
解释：从 [1, 1, 1] 开始
[1, 1, 1], 和为 3 ，选择下标 1
[1, 3, 1], 和为 5， 选择下标 2
[1, 3, 5], 和为 9， 选择下标 0
[9, 3, 5] 完成
示例 2：

输入：target = [1,1,1,2]
输出：false
解释：不可能从 [1,1,1,1] 出发构造目标数组。
示例 3：

输入：target = [8,5]    [1 , 1]
输出：true

提示：

N == target.length
1 <= target.length <= 5 * 10^4
1 <= target[i] <= 10^9

Python3+逆序+模拟+贪心

设数组长度为n则n为1时target必定只能为[1]
当n大于1时考虑逆序模拟依次将最大的数减到小于等于1为止
设当前target和为s最大值为a第二大值为b
逆向模拟减法操作首先需要将a减小到小于等于b
将数组分为和为s-a(含第二大值b)与a两部分
第一次减后数组变为s-a和2a-s假设此时2a-s>=b依旧成立
第二次减后数组变为s-a和3a-2s由此可知第k次减后数组为s-a与(k+1)a-ks
则若使(k+1)a-ks<=b成立有k>=(a-b)/(s-a)而k是整数因此向上取整
这个过程当中若是出现k为0的情况此时有a==b若b!=1则必然不可能成立
操作结束条件为target的最大值小于等于1
因为是辗转相减所以最大值一定会不断减小
最后检查得到的target是否符合要求
*/
func isPossible(target []int) bool {
	sort.Ints(target)
	sum := 0
	hp := new(myHeap)
	for _, v := range target {
		sum += v
		heap.Push(hp, v)
	}

	//x  s  t   (s-t)+x = t ==> x = 2*t-s

	for hp.Len() > 0 && (*hp)[0] != 1 && sum-(*hp)[0] != 1 {
		//最大值
		tmp := heap.Pop(hp).(int)
		//剩余和
		sum -= tmp

		if sum < 1 || tmp <= sum {
			return false
		}
		tmp %= sum
		sum += tmp
		if tmp == 0 {
			return false
		}
		heap.Push(hp, tmp)
	}
	return true
}

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
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
	*h = append(*h, v.(int))
}

// 8 5
//  1 1
//  1 2
//  3 2
//  3 5
//  8 5

func main() {
	fmt.Println(isPossible([]int{8, 5}))
	fmt.Println(isPossible([]int{9, 3, 5}))
}
