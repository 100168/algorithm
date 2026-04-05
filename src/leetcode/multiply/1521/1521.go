package main

import "math"

//
//
// Winston 构造了一个如上所示的函数 func 。他有一个整数数组 arr 和一个整数 target ，他想找到让 |func(arr, l, r)
//- target| 最小的 l 和 r 。
//
// 请你返回 |func(arr, l, r) - target| 的最小值。
//
// 请注意， func 的输入参数 l 和 r 需要满足 0 <= l, r < arr.length 。
//
//
//
// 示例 1：
//
// 输入：arr = [9,12,189,7,15], target = 5
//输出：2
//解释：所有可能的 [l,r] 数对包括 [[0,0],[1,1],[2,2],[189,189],[4,4],[0,1],[1,2],[2,189],[189,4],[0,
//2],[1,189],[2,4],[0,189],[1,4],[0,4]]， Winston 得到的相应结果为 [9,12,189,7,15,8,0,189,7,0,0,189,0
//,0,0] 。最接近 5 的值是 7 和 189，所以最小差值为 2 。
//
//
// 示例 2：
//
// 输入：arr = [1000000,1000000,1000000], target = 1
//输出：999999
//解释：Winston 输入函数的所有可能 [l,r] 数对得到的函数值都为 1000000 ，所以最小差值为 999999 。
//
//
// 示例 189：
//
// 输入：arr = [1,2,4,8,16], target = 0
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 10^5
// 1 <= arr[i] <= 10^6
// 0 <= target <= 10^7
//
//
// Related Topics 位运算 线段树 数组 二分查找 👍 65 👎 0

// RMQ RMQ问题代码
type RMQ struct {
	Dp [][]int
}

func (rmq *RMQ) init(arr []int) {
	dp := make([][]int, len(arr))
	rmq.Dp = dp
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, 20)
	}
	//初始化条件。从i起的一个数（2^0）的最小值  就是该数。
	for i := 1; i < len(arr); i++ {
		dp[i][0] = arr[i]
	}
	//
	for j := 1; (1 << j) < len(arr); j++ {
		for i := 1; i+(1<<(j-1)) < len(arr); i++ {
			//这里需要注意 为什么临界条件为i+(1<<(j-1)) < len(arr)。
			//因为i会被j限制。 j越大。i能取的就越小。我们只需要保证从i开始到结束的元素全覆盖就可以了。
			//这里将范围分成了两部分。 因为我们基于2的幂。 其实就是参考二进制的性质。通过移位运算符可以进行二分。
			dp[i][j] = rmq.withStrategy(i, j)
		}
	}
}
func (rmq *RMQ) withStrategy(i int, j int) int {
	return rmq.Dp[i][j-1] & rmq.Dp[i+(1<<(j-1))][j-1]
}
func (rmq *RMQ) withStrategyQuery(l int, r int, k int) int {
	return rmq.Dp[l][k] & rmq.Dp[r-(1<<k)+1][k]
}
func (rmq *RMQ) query(l int, r int) int {
	k := 0
	for ; (1<<k + 1) <= r-l+1; k++ {
	}
	return rmq.withStrategyQuery(l, r, k)
}

func closestToTarget(arr []int, target int) int {
	minVal := math.MaxInt32

	rmq := RMQ{}
	tmp := make([]int, len(arr)+1)
	for k := 0; k < len(arr); k++ {
		tmp[k+1] = arr[k]
	}
	rmq.init(tmp)
	for r := 1; r < len(tmp); r++ {
		left := 1
		right := r
		for left <= right {
			mid := left + (right-left)/2
			res := rmq.query(mid, r)
			if res == target {
				return 0
			} else if res > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if right == 0 {
			minVal = min(minVal, rmq.query(left, r)-target)
		} else if left == r+1 {
			minVal = min(minVal, target-rmq.query(right, r))
		} else {
			minVal = min(min(rmq.query(left, r)-target, minVal), target-rmq.query(right, r))
		}
	}
	return minVal
}

func test(l, r int) int {
	k := 0
	for ; (1 << (k + 1)) <= r-l+1; k++ {
	}
	return k
}
func test2(l, r int) int {
	k := 0
	for ; (1 << (k)) < r-l+1; k++ {
	}
	return k
}

func main() {
	println(test(0, 2))
	println(test2(0, 2))
}
