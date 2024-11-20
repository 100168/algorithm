package main

import "slices"

/**
给你一个整数 mountainHeight 表示山的高度。

同时给你一个整数数组 workerTimes，表示工人们的工作时间（单位：秒）。

工人们需要 同时 进行工作以 降低 山的高度。对于工人 i :

山的高度降低 x，需要花费 workerTimes[i] + workerTimes[i] * 2 + ... + workerTimes[i] * x 秒。例如：
山的高度降低 1，需要 workerTimes[i] 秒。
山的高度降低 2，需要 workerTimes[i] + workerTimes[i] * 2 秒，依此类推。
返回一个整数，表示工人们使山的高度降低到 0 所需的 最少 秒数。



示例 1：

输入： mountainHeight = 4, workerTimes = [2,1,1]

输出： 3

解释：

将山的高度降低到 0 的一种方式是：

工人 0 将高度降低 1，花费 workerTimes[0] = 2 秒。
工人 1 将高度降低 2，花费 workerTimes[1] + workerTimes[1] * 2 = 3 秒。
工人 2 将高度降低 1，花费 workerTimes[2] = 1 秒。
因为工人同时工作，所需的最少时间为 max(2, 3, 1) = 3 秒。

*/

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {

	//minV := slices.Min(workerTimes)
	maxV := slices.Max(workerTimes)

	l, r := 1, (1+mountainHeight)*mountainHeight/2*maxV

	check := func(t int) bool {
		cnt := 0
		for _, v := range workerTimes {

			if v > t {
				continue
			}
			tt := t * 2

			l, r := 1, t/v+1
			for l <= r {
				m := (r + l) / 2
				if m*(m+1) <= tt/v {
					l = m + 1
				} else {
					r = m - 1
				}
			}
			cnt += r
		}
		return cnt >= mountainHeight
	}

	for l <= r {
		m := (r + l) / 2
		if check(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return int64(l)
}
