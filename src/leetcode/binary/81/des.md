[二分搜索]()

1. 怎么找到单调性？
2. 怎么处理重复数据？

思路 : 尝试找规律，随机找一个数，观察这个数的左右两侧。发现一定会有一侧是单调递增的。

1. 定义两个指针l = 0，r = len(num) - 1
2. 中间值 m = l+(r-l)/2
3. 如果 nums[l] == nums[r] 则l++,向右找，为啥不r--?因为还不知道是否 target == nums[r]所以r也需要查找
4. 如果 nums[l] <= nums[m] 说明左边是递增的，此时：
    - 如果 nums[l] <= target<nums[m] 说明在左边区间内 r = m - 1
    - 否则往右边区间找 l = m + 1
5. 否则说明右边是递增的 
   - 如果 nums[m]<target<=nums[r] 说明在右边区间内 l = m + 1
   - 否则往左边区间找 r = m -1 

```go
package main

func search(nums []int, target int) bool {

   n := len(nums)
   l := 0
   r := n - 1

   for l <= r {
      m := l + (r-l)/2
      if nums[m] == target {
         return true
      }
      //相等说明有重复数据，一直往右边找
      if nums[l] == nums[r] {
         l++
         continue
      }
      if nums[l] <= nums[m] {
         if target >= nums[l] && target < nums[m] {
            r = m - 1
         } else {
            l = m + 1
         }
      } else {
         if target > nums[m] && target <= nums[n-1] {
            l = m + 1
         } else {
            r = m - 1
         }
      }

   }
   return false
}

```
