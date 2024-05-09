package main

import "fmt"

/*
*给你一个整数数组 arr ，请你删除一个子数组（可以为空），使得 arr 中剩下的元素是 非递减 的。

一个子数组指的是原数组中连续的一个子序列。

请你返回满足题目要求的最短子数组的长度。
*/
func findLengthOfShortestSubarray(arr []int) int {

	n := len(arr)

	left := make([]bool, n)
	right := make([]bool, n)

	mostLeft := 0
	mostRight := n - 1

	left[0] = true
	right[n-1] = true
	for i := 1; i < n; i++ {
		if arr[i] >= arr[i-1] && left[i-1] {
			left[i] = true
			mostLeft = i
		}
	}

	for i := n - 2; i >= 0; i-- {
		if arr[i] <= arr[i+1] && right[i+1] {
			right[i] = true
			mostRight = i
		}
	}

	if left[n-1] {
		return 0
	}
	l := 0

	//枚举右端点
	ans := min(n-mostLeft-1, mostRight)
	for r := 0; r < n; r++ {
		//左端点有序并且右端点有序并且值小于右端点
		for l < r && left[l] && right[r] && arr[l] <= arr[r] {
			ans = min(ans, max(r-l-1, 0))
			l++
		}
	}
	return ans
}

/**
class Solution {
    public int findLengthOfShortestSubarray(int[] arr) {
        int n = arr.length;
        int i = 1, j = n-1;
        while (i<n && arr[i-1]<=arr[i]) ++i;
        if (i == n) return 0; // arr已经有序
        while (j-1>=0 && arr[j-1]<=arr[j]) --j;
        int ans = j; // 只保留右边
        int l = 0; // l指针初始为left左端点
        int r = j; // r指针初始为right左端点
        while (l < i && r < n) { //两个指针均不越界
            if (arr[l] <= arr[r]) {
                ans = Math.min(ans, r-l-1); // 取窗口最小值
                l++; // 滑动窗口l指针
            } else {
                r++; // 滑动窗口r指针
            }
        }
         ans = Math.min(ans, n-i); // 只保留左边，再取最小值
        return ans;
    }
}
*/

func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{10, 13, 17, 21, 15, 15, 9, 17, 22, 22, 13}))
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 5, 4, 3, 4, 5}))
	fmt.Println(findLengthOfShortestSubarray([]int{5, 4, 3, 2, 1, 3, 4, 5, 6, 4, 3, 494, 304, 3, 4, 3, 4, 5, 6, 3, 4, 5}))
}
