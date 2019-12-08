package main

import "fmt"

func longestPalindrome(s string) string {
	// f[i,j] = f[i+1, j-1] && a[i] == a[j]

	l := len(s)
	if l <= 0 {
		return ""
	}
	f := make([][]bool, l)
	for i := 0; i < l; i++ {
		f[i] = make([]bool, l)
	}

	for i := 0; i < l; i++ {
		for j := 0; j <= i; j++ {
			f[i][j] = true
		}
	}
	res := s[:1]
	max := 1
	for i := l - 2; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			f[i][j] = f[i+1][j-1] && s[i] == s[j]
			if f[i][j] && j-i+1 > max {
				max = j - i + 1
				res = s[i : j+1]
			}
		}
	}

	return res
}

func main() {
	ret := longestPalindrome("babad")
	fmt.Println(ret)
	ret = longestPalindrome("cbbd")
	fmt.Println(ret)
}
