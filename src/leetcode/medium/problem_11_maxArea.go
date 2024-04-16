package main

func maxArea(height []int) int {

	n := len(height)
	l := 0
	r := n - 1
	maxArea := 0

	for l < r {
		if height[l] > height[r] {
			if height[r]*(r-l+1) > maxArea {
				maxArea = height[r] * (r - l + 1)
			}
			r--
		} else {
			if height[l]*(r-l+1) > maxArea {
				maxArea = height[l] * (r - l + 1)
			}
			l++
		}
	}

	return maxArea
}
func main() {

}
