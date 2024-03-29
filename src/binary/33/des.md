[二分搜索]()  
随机选一个数之后，这个数的左边或者右边必有一个区间是单调递增的

1. 定义左右指针l,r
2. 中间值 m = (l+r)/2 并拿中间值跟最左边的数比教判断单调性
    - 如果 nums[m] <= nums[0] 左边单调递增 右边不确定
        - 如果 nums[0]<=target<nums[m] 说明在左区间 ==> r = m -1
        - 否则 说明在右边区间 ==> l = m+1
    - 否则 右边单调递增，左边不确定，此时再判断target是否在右边区间
        - 如果 nums[m]<target<=nums[len(nums)-1] 说明在左边区间 ==> l = m+1
        - 否则 说明在右边区间 ==> r = m-1





```go
package main

func search(nums []int, target int) int {

	l := 0
	r := len(nums) - 1
	n := len(nums) - 1

	// Input: nums = [4,5,6,7,0,1,2], target = 3
	//这里用<=是因为每次查找都不包括中间的数
	for l <= r {
		m := (r + l) / 2
		if nums[m] == target {
			return m
		}

		//1.如果中间数大于等于最左边的数，说明左半区间是递增的
		//2.否则说明右半区间是递增的
		if nums[m] >= nums[0] {
			//1.目标数比当前数小，并且大于等于最左边数，往左半区间找
			//2.否则说明当前数在右半区间
			if nums[0] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			//1.到这说明右边是递增的
			//2.如果目标数大于当前数，并且小于等于最右边数，说明要往右半区间找
			//3.否则往左半区间找
			if target > nums[m] && target <= nums[n] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	//说明没找到，返回-1

	return -1

}

```






