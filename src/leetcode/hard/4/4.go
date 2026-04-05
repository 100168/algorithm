package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1 := len(nums1)
	l2 := len(nums2)
	//true 为偶数
	even := (l1+l2)&1 == 0

	mid := (l1 + l2) >> 1
	if l1 > 0 && l2 > 0 {
		find1 := findK(nums1, nums2, mid+1)
		if even {
			find2 := findK(nums1, nums2, mid)
			return (float64(find1) + float64(find2)) * 0.5
		}
		return float64(find1)
	} else if l1 != 0 {
		if even {
			return (float64(nums1[mid-1]) + float64(nums1[mid])) * 0.5
		}
		return float64(nums1[mid-1])
	} else {
		if even {
			return (float64(nums2[mid-1]) + float64(nums2[mid])) * 0.5
		}
		return float64(nums2[mid-1])
	}
}

func findK(nums1 []int, nums2 []int, k int) int {

	var longest []int
	var shortest []int
	l1 := len(nums1)
	l2 := len(nums2)

	if l1 > l2 {
		longest = nums1
		shortest = nums2
	} else {
		longest = nums2
		shortest = nums1
	}
	s := len(shortest)
	l := len(longest)

	if k <= s {
		return getUpMedium(shortest, 0, k-1, longest, 0, k-1)
	} else if k > l {

		if longest[k-s-1] >= shortest[s-1] {
			return longest[k-s-1]
		}
		if shortest[k-l-1] >= longest[l-1] {
			return shortest[k-l-1]
		}
		//?为啥是这样取？因为前k-l个数不要
		return getUpMedium(shortest, k-l, s-1, longest, k-s, l-1)
	}

	if longest[k-s-1] > shortest[s-1] {
		return longest[k-s-1]
	}
	return getUpMedium(shortest, 0, s-1, longest, k-s, k-1)
}

func getUpMedium(shortest []int, s1 int, e1 int, longest []int, s2 int, e2 int) int {
	for s1 < e1 {
		mid1 := (s1 + e1) >> 1
		mid2 := (s2 + e2) >> 1
		offset := ((e1 - s1 + 1) & 1) ^ 1
		if shortest[mid1] == longest[mid2] {
			return shortest[mid1]
		} else if shortest[mid1] > longest[mid2] {
			e1 = mid1
			s2 = mid2 + offset
		} else {
			s1 = mid1 + offset
			e2 = mid2
		}
	}
	if shortest[s1] > longest[s2] {
		return longest[s2]
	}
	return shortest[s1]
}

func main() {
	nums1 := []int{1}
	nums2 := []int{1}
	val := findMedianSortedArrays(nums1, nums2)
	fmt.Println(val)
}
