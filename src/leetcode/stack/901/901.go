package main

/**
设计一个算法收集某些股票的每日报价，并返回该股票当日价格的 跨度 。

当日股票价格的 跨度 被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。

例如，如果未来 7 天股票的价格是 [100,80,60,70,60,75,85]，那么股票跨度将是 [1,1,1,2,1,4,6] 。

实现 StockSpanner 类：

StockSpanner() 初始化类对象。
int next(int price) 给出今天的股价 price ，返回该股票当日价格的 跨度 。*/

type StockSpanner struct {
	nums []int
	st   []int
}

func Constructor() StockSpanner {

	ss := new(StockSpanner)
	ss.st = make([]int, 0)
	ss.nums = make([]int, 0)
	return *ss
}

func (ss *StockSpanner) Next(price int) int {
	ss.nums = append(ss.nums, price)
	nums := ss.nums
	cur := ss.st
	for len(cur) > 0 && nums[cur[len(cur)-1]] <= price {
		cur = cur[:len(cur)-1]
	}
	mostLeft := 0
	if len(cur) > 0 {
		mostLeft = cur[len(cur)-1] + 1
	}
	cur = append(cur, len(nums)-1)
	ss.st = cur
	return len(ss.nums) - mostLeft
}
