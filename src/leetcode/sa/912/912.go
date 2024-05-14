package main

import (
	"fmt"
	"math/rand"
	"slices"
)

// 基数排序
func sortArray(nums []int) []int {
	n := len(nums)

	for i := range nums {
		nums[i] += 50000
	}
	maxVal := slices.Max(nums)

	for exp := 1; exp < maxVal; exp *= 10 {
		cnt := make([]int, 10)
		//sa[i]表示排名为i的数字在数组中的位置，也就是排名为i的是哪个数字
		sa := make([]int, n)
		for _, v := range nums {
			cur := (v / exp) % 10
			cnt[cur]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}

		for i := n - 1; i >= 0; i-- {
			cur := (nums[i] / exp) % 10
			sa[cnt[cur]-1] = nums[i]
			cnt[cur]--
		}
		nums = sa
	}

	for i := range nums {
		nums[i] -= 50000
	}
	return nums

}

/*
func sortArray(nums []int) []int {

     n:=len(nums)


     maxVal:=nums[0]

     for i:=range nums{
        nums[i]+=50000
        maxVal = max(maxVal,nums[i])
     }

     for exp:=1;exp<=maxVal;exp*=10{

        cnt:=make([]int,10)
        sa:=make([]int,n)

        for i:=0;i<n;i++{
            cur :=nums[i]/exp%10
            cnt[cur]++
        }
        for i:=1;i<10;i++{
            cnt[i] +=cnt[i-1]
        }
        for i:=n-1;i>=0;i--{
            cur:=nums[i]/exp%10
            sa[cnt[cur]-1] = nums[i]
            cnt[cur]--

        }
        copy(nums,sa)
     }
     for i:=range nums{
        nums[i]-=50000
     }
     return nums
}*/

func sortArray2(nums []int) []int {

	quicklySort(0, len(nums)-1, nums)
	return nums
}

// 快速排序
func quicklySort(l, r int, nums []int) {
	if l >= r {
		return
	}
	m := partition(l, r, nums)
	quicklySort(l, m-1, nums)
	quicklySort(m+1, r, nums)
}
func partition(left int, right int, nums []int) int {

	c := rand.Intn(right-left+1) + left
	flag := nums[c]
	nums[c], nums[right] = nums[right], nums[c]
	index := left - 1
	for i := left; i < right; i++ {
		if nums[i] <= flag {
			index++
			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	index++
	nums[index], nums[right] = nums[right], nums[index]
	return index

}

func main() {
	fmt.Println(sortArray([]int{-2, 3, -5}))
	fmt.Println(sortArray2([]int{-2, 3, -5}))
}
