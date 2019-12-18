package main

import "fmt"
import "sort"

func makeSet(nums []int) map[int]int {
	var ret map[int]int = map[int]int{}
	for _, v := range nums {
		ret[v] ++
	}

	return ret
}

func threeSum(nums []int) [][]int {
	var ret = [][]int{}
	sort.Ints(nums)

	l := len(nums)
	if l <= 0 {
		return ret
	}
	set := makeSet(nums)
	var preI, preJ int = nums[0] - 1, nums[0] - 1
	for i := 0; i < l-2; i ++ {
		if preI == nums[i]{
			continue
		}
		preI = nums[i]
		preJ = nums[0] - 1
		set[nums[i]] --
		for j := i + 1; j < l-1; j ++ {
			if preJ == nums[j] {
				continue
			}
			preJ = nums[j]
			set[nums[j]] --

			k := 0 - nums[i] - nums[j]
			if (set[k] >= 1 && k >= nums[i] && k >= nums[j]) {
				ret = append(ret, []int{nums[i], nums[j], k})
			}
			set[nums[i]] ++
			set[nums[j]] ++
		}
	}
	return ret
}

func main() {
	input := [][]int{[]int{-1, 0, 1, 2, -1, -4}, []int{1, -1, -1, 0}}
	for _, v := range input {
		fmt.Println(threeSum(v))
	}
}