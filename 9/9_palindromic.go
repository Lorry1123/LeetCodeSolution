package main

import "fmt"

func isPalindrome(x int) bool {
	// 要求不转成字符串或数组....明明时间复杂度都一样....
	// 限制不了只能提要求....leetcode 的无能狂怒
	if x < 0 {
		return false
	}

	d := 1
	tmp := x
	for tmp/10 > 0 {
		d *= 10
		tmp /= 10
	}

	for x > 0 && d > 0 {
		if x/d != x%10 {
			return false
		}

		x = x - x/d*d
		x /= 10
		d /= 100
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-123))
	fmt.Println(isPalindrome(33))
}
