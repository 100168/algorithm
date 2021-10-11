package main

/*给你一个正整数数组arr，请你计算所有可能的奇数长度子数组的和。

子数组 定义为原数组中的一个连续子序列。

请你返回 arr中 所有奇数长度子数组的和 。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	sumPre := make([]int, n)
	sumPre[0] = arr[0]
	for i := 1; i < n; i++ {
		sumPre[i] = sumPre[i-1] + arr[i]
	}
	sum := 0

	for k := 1; k < n; k = 2*k + 1 {
		for i := k - 1; i < n; i++ {
			if i == k-1 {
				sum += sumPre[i]
			} else {
				sum += sumPre[i] - sumPre[i-k]
			}
		}
	}
	return sum
}
func main() {

}
