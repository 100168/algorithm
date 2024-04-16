package main

//给你两个正整数 n 和 limit 。
//
// 请你将 n 颗糖果分给 3 位小朋友，确保没有任何小朋友得到超过 limit 颗糖果，请你返回满足此条件下的 总方案数 。
//
//
//
// 示例 1：
//
//
//输入：n = 5, limit = 2
//输出：3
//解释：总共有 3 种方法分配 5 颗糖果，且每位小朋友的糖果数不超过 2 ：(1, 2, 2) ，(2, 1, 2) 和 (2, 2, 1) 。
//
//
// 示例 2：
//
//
//输入：n = 3, limit = 3
//输出：10
//解释：总共有 10 种方法分配 3 颗糖果，且每位小朋友的糖果数不超过 3 ：(0, 0, 3) ，(0, 1, 2) ，(0, 2, 1) ，(0, 3,
// 0) ，(1, 0, 2) ，(1, 1, 1) ，(1, 2, 0) ，(2, 0, 1) ，(2, 1, 0) 和 (3, 0, 0) 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁶
// 1 <= limit <= 10⁶
//
//
// Related Topics 数学 组合数学 枚举 👍 12 👎 0

func distributeCandies(n int, limit int) int64 {

	ans := int64(0)
	//先枚举第一个数
	for x := 0; x <= min(limit, n); x++ {
		//for y:=max(n-x-limit,0);y<=min(n-x,limit);y++{
		//
		//}
		//然后确定第二个数范围
		//计算范围内的值
		//left := max(n-x-limit, 0)
		//right := min(n-x, limit)

		ans += int64(max(min(n-x, limit)-max(n-x-limit, 0)+1, 0))
	}

	//n-x-limit
	return ans
}

func main() {
	println(distributeCandies(4, 1))
}
