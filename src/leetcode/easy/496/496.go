package main

/*
给你两个 没有重复元素 的数组nums1 和nums2，其中nums1是nums2的子集。

请你找出 nums1中每个元素在nums2中的下一个比其大的值。

nums1中数字x的下一个更大元素是指x在nums2中对应位置的右边的第一个比x大的元素。如果不存在，对应位置输出 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
使用了单调栈
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	paris := make(map[int]int)
	m := len(nums2)
	var stack []int
	for i := 0; i < m; i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			paris[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	for i := range nums1 {
		nums1[i] = paris[nums1[i]]
		if nums1[i] == 0 {
			nums1[i] = -1
		}
	}
	return nums1

}
func main() {

}
