package main

import "fmt"

func maxArea(height []int) int {
	l := len(height)
	ret := 0
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			tmp := 0
			if height[i] > height[j] {
				tmp = height[j] * (j - i)
			} else {
				tmp = height[i] * (j - i)
			}
			if tmp > ret {
				ret = tmp
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
