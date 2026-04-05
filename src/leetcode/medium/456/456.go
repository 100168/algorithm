package main

import (
	"fmt"
	"math"
)

/*
*给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，
并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。

如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/132-pattern
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

使用单调栈，从右往左遍历，*
*/
func find132pattern(nums []int) bool {
	n := len(nums)
	//123中的3
	candidateK := []int{nums[n-1]}

	//123中的2
	maxK := math.MinInt64

	/*//单调递增栈：假如有一个元素要入栈并且它比部分的栈内元素大，那么将部分栈内所有元素取出，
	//直到比入栈元素大的元素，入栈，也就是必须要保持栈顶到栈底元素是单调递增的
	//从右往左迭代序列，假如数值一直在变小或相等，就入栈，直到找到比栈顶元素大的元素后，
	//将栈顶元素取出并设为max_k，之后找到比max_k小的元素就返回True*/
	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		//为什么要用while,是因为这是一个单调栈，找到所有比nums[i]小的元素并取出
		//candidate_k在pop之后更新了，所以会再执行一遍while循环看是否符合条件
		//这样就找到了比nums[i]小但是差距又最小的值，使得max_k仅比nums[i]小但是又越大越好
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}

	return false
}

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4}

	pattern2 := find132pattern(a)
	fmt.Println(pattern2)
}
