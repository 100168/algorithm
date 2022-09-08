package main

/**给定两个以 升序排列 的整数数组 nums1 和 nums2,以及一个整数 k。

定义一对值(u,v)，其中第一个元素来自nums1，第二个元素来自 nums2。

请找到和最小的 k个数对(u1,v1), (u2,v2) ... (uk,vk)。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	l1 := len(nums1)
	l2 := len(nums2)
	index1 := 0
	index2 := 0

	ans := make([][]int, 2)
	ans = append(ans, []int{nums1[0], nums2[0]})
	k--

	for index1 < l1 && index2 < l2 {
		//res := make([]int, 2)
		if k == 0 {
			return ans
		}

		//if index1+1 < l1 && nums1[index1] < {
		//
		//}

	}
	return ans
}
