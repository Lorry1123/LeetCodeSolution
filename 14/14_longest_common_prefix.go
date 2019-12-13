package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	minLength := 0x7f7f7f7f
	for _, v := range strs {
		if len(v) < minLength {
			minLength = len(v)
		}
	}

	ret := ""
	for i := 0; i < minLength; i ++ {
		tmp := strs[0][i]
		for _, v := range strs {
			if v[i] == tmp {
				continue
			}
			return ret
		}
		ret += string(tmp)
	}
	return ret
}

func main() {
	input := [][]string{[]string{"flower","flow","flight"}, []string{"dog","racecar","car"}}
	for _, v := range input {
		fmt.Println(v)
		fmt.Println(longestCommonPrefix(v))
	}
}