package main

/*
这里有 n 个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。

请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
*/
func corpFlightBookings(bookings [][]int, n int) []int {

	diff := make([]int, n+2)
	ans := make([]int, n)
	for _, v := range bookings {
		diff[v[0]] += v[2]
		diff[v[1]+1] -= v[2]
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += diff[i]
		ans[i-1] = sum
	}
	return ans

}
