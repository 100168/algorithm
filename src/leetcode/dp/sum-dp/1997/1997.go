package main

/*
你需要访问 n 个房间，房间从 0 到 n - 1 编号。同时，每一天都有一个日期编号，从 0 开始，依天数递增。你每天都会访问一个房间。

最开始的第 0 天，你访问 0 号房间。给你一个长度为 n 且 下标从 0 开始 的数组 nextVisit 。在接下来的几天中，你访问房间的 次序 将根
据下面的 规则 决定：


假设某一天，你访问 i 号房间。
如果算上本次访问，访问 i 号房间的次数为 奇数 ，那么 第二天 需要访问 nextVisit[i] 所指定的房间，其中 0 <= nextVisit[
i] <= i 。
如果算上本次访问，访问 i 号房间的次数为 偶数 ，那么 第二天 需要访问 (i + 1) mod n 号房间。


请返回你访问完所有房间的第一天的日期编号。题目数据保证总是存在这样的一天。由于答案可能很大，返回对 10⁹ + 7 取余后的结果。



示例 1：


输入：nextVisit = [0,0]
输出：2
解释：
- 第 0 天，你访问房间 0 。访问 0 号房间的总次数为 1 ，次数为奇数。
下一天你需要访问房间的编号是 nextVisit[0] = 0
- 第 1 天，你访问房间 0 。访问 0 号房间的总次数为 2 ，次数为偶数。
下一天你需要访问房间的编号是 (0 + 1) mod 2 = 1
- 第 2 天，你访问房间 1 。这是你第一次完成访问所有房间的那天。


示例 2：


输入：nextVisit = [0,0,2]
输出：6
解释：
你每天访问房间的次序是 [0,0,1,0,0,1,2,...] 。
第 6 天是你访问完所有房间的第一天。


示例 3：


输入：nextVisit = [1,0,2,0]
 0,1,0,1,2
输出：6
解释：
你每天访问房间的次序是 [0,0,1,1,2,2,3,...] 。
第 6 天是你访问完所有房间的第一天。


提示：
n == nextVisit.length
2 <= n <= 10⁵
0 <= nextVisit[i] <= i

odd[i] := even[i-1]+1

even[i]:= s(even[j]+even[j-1]+2)

*/

func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)
	mod := 1_000_000_007
	s := make([]int, n)

	//hoe to think?
	//第一次到i访问次数是奇数，只能往前跳，i之前的的位置访问的次数一定是偶数 f[i] = f[]
	//第二次到i是偶数，可以往右跳

	//f[0]+f[1]+⋯+f[n−2]+1=s[n−1]+1
	//即定义 f[i] 表示从「访问到房间 i 且次数为奇数」到「访问到房间 i 且次数为偶数」所需要的天数。
	for i, j := range nextVisit[:n-1] {
		s[i+1] = (s[i]*2 - s[j] + 2 + mod) % mod
	}
	return s[n-1]
}

func firstDayBeenInAllRooms2(nextVisit []int) int {
	const mod = 1_000_000_007
	n := len(nextVisit)
	s := make([]int, n)
	for i, j := range nextVisit[:n-1] {
		s[i+1] = (s[i]*2 - s[j] + 2 + mod) % mod // + mod 避免算出负数
	}
	return s[n-1]
}
