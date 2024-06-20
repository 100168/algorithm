package main

import "fmt"

/*
*
给你一个下标从 0 开始的整数数组 nums 。如果下标对 i、j 满足 0 ≤ i < j < nums.length ，
如果 nums[i] 的 第一个数字 和 nums[j] 的 最后一个数字 互质 ，则认为 nums[i] 和 nums[j] 是一组 美丽下标对 。

返回 nums 中 美丽下标对 的总数目。
*/
func countBeautifulPairs(nums []int) int {

    ans := 0
    cntLast := make([]int, 10)

    for i := range nums {
        cur := nums[i]
        cntLast[cur%10]++
    }

    for i := range nums {
        cntLast[nums[i]%10]--

        cur := nums[i]
        for cur >= 10 {
            cur /= 10
        }
        for j, v := range cntLast {
            if gcd(cur, j) == 1 {
                ans += v
            }
        }
    }
    return ans

}

func gcd(a, b int) int {

    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    fmt.Println(countBeautifulPairs([]int{31, 25, 72, 79, 74}))
}
