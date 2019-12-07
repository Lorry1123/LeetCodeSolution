package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var flag [256]int
	st := 0
	ed := 0
	l := len(s)

	ans := 0

	for ed < l {
		if flag[s[ed]] > 0 {
			flag[s[st]] --
			st ++
		} else {
			flag[s[ed]] ++
			ed ++
		}

		if ed - st > ans {
			ans = ed - st
		}
	}

	return ans
}

func main() {
	ret := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(ret)

	ret = lengthOfLongestSubstring("bbbbbb")
	fmt.Println(ret)

	ret = lengthOfLongestSubstring("pwwkew")
	fmt.Println(ret)

	ret = lengthOfLongestSubstring("  ")
	fmt.Println(ret)
}
