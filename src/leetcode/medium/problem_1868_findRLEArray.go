package main

import "fmt"

/*
行程编码（Run-length encoding）是一种压缩算法，能让一个含有许多段连续重复数字的整数类型数组nums以一个（通常更小的）二维数组encoded表示。每个 encoded[i] = [vali, freqi] 表示 nums 中第 i 段重复数字，其中 vali 是该段重复数字，重复了 freqi 次。

例如，nums = [1,1,1,2,2,2,2,2]可表示称行程编码数组encoded = [[1,3],[2,5]]。对此数组的另一种读法是“三个1，后面有五个2”。
两个行程编码数组 encoded1和encoded2 的积可以按下列步骤计算：

将encoded1和encoded2分别扩展成完整数组nums1和nums2 。
创建一个新的数组prodNums，长度为nums1.length并设prodNums[i] = nums1[i] * nums2[i]。
将prodNums压缩成一个行程编码数组并返回之。
给定两个行程编码数组encoded1 和 encoded2 ，分别表示完整数组nums1 和 nums2 。nums1 和 nums2 的长度相同。 每一个 encoded1[i] = [vali, freqi] 表示 nums1 中的第 i 段，每一个 encoded2[j] = [valj, freqj] 表示 nums2 中的第 j 段。

返回 encoded1 和 encoded2 的乘积。

注：行程编码数组需压缩成可能的最小长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/product-of-two-run-length-encoded-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findRLEArray(encoded1 [][]int, encoded2 [][]int) (ans [][]int) {
	i, j, n1, n2 := 0, 0, len(encoded1), len(encoded2)
	for i < n1 && j < n2 {
		//找最小
		minLen := minLen(encoded1[i][1], encoded2[j][1])
		mul := encoded1[i][0] * encoded2[j][0]
		encoded1[i][1] -= minLen
		encoded2[j][1] -= minLen
		cur := []int{mul, minLen}
		if len(ans) > 0 && ans[len(ans)-1][0] == mul { //to add up
			ans[len(ans)-1][1] += minLen
		} else {
			ans = append(ans, cur)
		}

		if encoded1[i][1] == 0 {
			i++
		}

		if encoded2[j][1] == 0 {
			j++
		}
	}

	return
}
func minLen(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

	/*[[1,3],[2,3]]
	[[6,3],[3,3]]*/
	a := [][]int{{1, 3}, {2, 3}}
	b := [][]int{{6, 3}, {3, 3}}

	array := findRLEArray(a, b)

	fmt.Println(array)
}
