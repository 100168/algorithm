package main

func longestConsecutive(nums []int) int {
	numsSet := map[int]bool{}
	for i := range nums {
		numsSet[nums[i]] = true
	}
	longest := 0
	for i := range numsSet {
		if !(numsSet[i-1]) {
			currentLong := 1
			currentNum := i + 1
			for numsSet[currentNum] {
				currentLong++
				currentNum++
			}
			if currentLong > longest {
				longest = currentLong
			}

		}
	}
	return longest
}
func main() {

}
