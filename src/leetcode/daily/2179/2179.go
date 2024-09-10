package main

import "fmt"

/*
*
给你两个下标从 0 开始且长度为 n 的整数数组 nums1 和 nums2 ，两者都是 [0, 1, ..., n - 1] 的 排列 。

好三元组 指的是 3 个 互不相同 的值，且它们在数组 nums1 和 nums2 中出现顺序保持一致。换句话说，如果我们将 pos1v 记为值 v 在 nums1 中出现的位置，pos2v 为值 v 在 nums2 中的位置，那么一个好三元组定义为 0 <= x, y, z <= n - 1 ，且 pos1x < pos1y < pos1z 和 pos2x < pos2y < pos2z 都成立的 (x, y, z) 。

请你返回好三元组的 总数目 。

示例 1：

输入：nums1 = [2,0,1,3], nums2 = [0,1,2,3]
输出：1
解释：
总共有 4 个三元组 (x,y,z) 满足 pos1x < pos1y < pos1z ，分别是 (2,0,1) ，(2,0,3) ，(2,1,3) 和 (0,1,3) 。
这些三元组中，只有 (0,1,3) 满足 pos2x < pos2y < pos2z 。所以只有 1 个好三元组。
示例 2：

输入：nums1 = [4,0,1,3,2], nums2 = [4,1,0,2,3]
输出：4
解释：总共有 4 个好三元组 (4,0,3) ，(4,0,2) ，(4,1,3) 和 (4,1,2) 。

n == nums1.length == nums2.length
3 <= n <= 105
0 <= nums1[i], nums2[i] <= n - 1
nums1 和 nums2 是 [0, 1, ..., n - 1] 的排列。
*/
func goodTriplets(nums1 []int, nums2 []int) int64 {

	n := len(nums1)

	indexMap := make(map[int]int)

	for i, v := range nums2 {
		indexMap[v] = i
	}

	bt := new(bitTree)

	bt.sum = make([]int, n+1)
	bt.l = n + 1
	nums := make([]int, n)

	cnt := 0

	for i := range nums {
		id := indexMap[nums1[i]] + 1
		//左边小于当前数个个数
		less := bt.query(id)
		//左边大于当前数的个数
		leftMore := i - less
		//大于当前数的个数
		restMore := n - id
		//右边大于单前数的个数 = 总的大于当前数个数-左边大于当前数的个数
		more := restMore - leftMore
		cnt += less * more
		//更新个数
		bt.update(id, 1)
	}
	return int64(cnt)

}

type bitTree struct {
	sum []int
	l   int
}

func lowBit(index int) int {
	return index & -index
}

func (bt bitTree) query(index int) int {

	ans := 0
	for index > 0 {
		ans += bt.sum[index]
		index -= lowBit(index)
	}
	return ans
}
func (bt bitTree) update(index int, val int) {
	for index < bt.l {
		bt.sum[index] += val
		index += lowBit(index)
	}
}

func main() {
	//3,1,2,4
	fmt.Println(goodTriplets([]int{2, 0, 1, 3}, []int{0, 1, 2, 3}))
	fmt.Println(goodTriplets([]int{4, 0, 1, 3, 2}, []int{4, 1, 0, 2, 3}))
}
