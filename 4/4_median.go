package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	len1, len2 := len(nums1), len(nums2)

	l := 0
	r := len1
	var maxl, minr int
	for l <= r {
		mid1 := (l + r) >> 1
		mid2 := (len1+len2+1)/2 - mid1

		if mid1 < len1 && nums2[mid2-1] > nums1[mid1] {
			l = mid1 + 1
		} else if mid1 > 0 && nums1[mid1-1] > nums2[mid2] {
			r = mid1 - 1
		} else {
			if mid1 == 0 {
				maxl = nums2[mid2-1]
			} else if mid2 == 0 {
				maxl = nums1[mid1-1]
			} else {
				maxl = max(nums1[mid1-1], nums2[mid2-1])
			}

			if (len1+len2)%2 == 1 {
				return float64(maxl)
			}

			if mid1 == len1 {
				minr = nums2[mid2]
			} else if mid2 == len2 {
				minr = nums1[mid1]
			} else {
				minr = min(nums1[mid1], nums2[mid2])
			}

			return float64(minr+maxl) / 2.0
		}
	}
	return -1
}

func main() {
	ret := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println(ret)
	ret = findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(ret)
	ret = findMedianSortedArrays([]int{1}, []int{1})
	fmt.Println(ret)
	ret = findMedianSortedArrays([]int{5, 6}, []int{1, 2, 3, 4, 7})
	fmt.Println(ret)
}
