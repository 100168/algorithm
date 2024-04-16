[思路]() 

1. 怎么找到单调性？ 

思路 : 用最后一个数跟中间数比较可以确定最小数在哪个区间

1. 定义两个指针l = 0，r = len(num) - 1
2. 中间值 m = l+(r-l)/2
3. 如果 nums[m] <= nums[r] 则r=m,向左找
4. 如果 nums[] <= nums[m] 说明左边是递增的，此时：
    - 如果 nums[l] <= target<nums[m] 说明在左边区间内 r = m - 1
    - 否则往右边区间找 l = m + 1
5. 否则说明右边是递增的
    - 如果 nums[m]<target<=nums[r] 说明在右边区间内 l = m + 1
    - 否则往左边区间找 r = m -1
```go
package main
func findMin(nums []int) int {

	n := len(nums)
	l := 0
	r := n - 1
	for l < r {
		m := l + (r-l)/2
		if nums[m] <= nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}
	
	return nums[l]
}
```