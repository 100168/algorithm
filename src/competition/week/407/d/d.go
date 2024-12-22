package main

import (
	"fmt"
)

/*
*

给你两个长度相同的正整数数组 nums 和 target。

在一次操作中，你可以选择 nums 的任何
子数组
，并将该子数组内的每个元素的值增加或减少 1。

返回使 nums 数组变为 target 数组所需的 最少 操作次数。

示例 1：

输入： nums = [3,5,1,2], target = [4,6,2,4]

输出： 2

解释：

执行以下操作可以使 nums 等于 target：
- nums[0..3] 增加 1，nums = [4,6,2,3]。
- nums[3..3] 增加 1，nums = [4,6,2,4]。

示例 2：

输入： nums = [1,3,2], target = [2,1,4]

输出： 5

解释：

执行以下操作可以使 nums 等于 target：
- nums[0..0] 增加 1，nums = [2,3,2]。
- nums[1..1] 减少 1，nums = [2,2,2]。
- nums[1..1] 减少 1，nums = [2,1,2]。
- nums[2..2] 增加 1，nums = [2,1,3]。
- nums[2..2] 增加 1，nums = [2,1,4]。

提示：

1 <= nums.length == target.length <= 105
1 <= nums[i], target[i] <= 108
*/
func minimumOperations(nums, target []int) int64 {
	s := target[0] - nums[0]
	ans := abs(s)
	for i := 1; i < len(nums); i++ {
		k := (target[i] - target[i-1]) - (nums[i] - nums[i-1])
		if k > 0 {
			if s >= 0 {
				ans += k
			} else {
				ans += max(k+s, 0)
			}
		} else {
			if s <= 0 {
				ans -= k
			} else {
				ans -= min(k+s, 0)
			}
		}
		s += k
	}
	return int64(ans)
}

/*
*
在 d=[1,−3,4,−2] 这个例子中，d[0]=1 和 d[2]=4 对应着 5 个 +1 操作。由于我们是从全 0 的差分数组开始的，所以一定有 5 个对应的 −1 操作，这正好是 d[1]=−3 和 d[3]=−2。

如果同一个 d[i] 上面既有 +1，又有 −1，那这肯定不如只有 +1 或者只有 −1 优。所以最优操作中，没有 +1 和 −1「互相抵消」的情况。

由于每次操作会产生一个 +1 和一个 −1，所以操作次数就等于所有正数 d[i] 之和。
*/
func minimumOperations2(nums, target []int) int64 {
	n := len(nums)
	ps := 0
	negs := 0
	if target[0]-nums[0] < 0 {
		negs += target[0] - nums[0]
	} else {
		ps += target[0] - nums[0]
	}
	for i := 1; i < n; i++ {
		cur := target[i] - nums[i] - (target[i-1] - nums[i-1])

		if cur > 0 {
			ps += cur
		} else {
			negs += cur
		}

	}

	return int64(max(-negs, ps))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	//fmt.Println(minimumOperations([]int{3, 5, 1, 2}, []int{4, 6, 2, 4}))
	fmt.Println(minimumOperations([]int{1, 3, 2}, []int{2, 1, 4}))
}
