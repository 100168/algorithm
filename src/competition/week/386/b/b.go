package main

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {

	ans := 0

	for i, v := range bottomLeft {

		x1, y1 := v[0], v[1]
		x2, y2 := topRight[i][0], topRight[i][1]

		for j := range bottomLeft {
			if i == j {
				continue
			}
			xx1, yy1 := bottomLeft[j][0], bottomLeft[j][1]
			xx2, yy2 := topRight[j][0], topRight[j][1]
			curY := max(min(yy2, y2)-max(y1, yy1), 0)
			curX := max(min(xx2, x2)-max(xx1, x1), 0)
			cur := min(curX, curY)
			ans = max(ans, cur)
		}
	}
	return int64(ans * ans)
}
