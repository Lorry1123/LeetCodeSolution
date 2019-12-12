package main

import "fmt"

func intToRoman(num int) string {
	words := map[int]string{
		1000: "M",
		500: "D",
		100: "C",
		50: "L",
		10: "X",
		5: "V",
		1: "I",
		0: "",
	}

	iter := []int{1000, 500, 100, 50, 10, 5, 1}
	p := 0
	ret := ""
	for num > 0 && p < 7{
		for num >= iter[p] {
			ret += words[iter[p]]
			num -= iter[p]
		}

		lower := p + 1
		if lower % 2 != 0 {
			lower ++
		}

		if lower >= 7 {
			p ++
			continue
		}

		if num >= iter[p] - iter[lower] {
			ret = ret + words[iter[lower]] + words[iter[p]]
			num -= iter[p] - iter[lower]
		}
		p ++
	}

	return ret
}


func main() {
	input := []int{3, 4, 9, 58, 1994}
	for _, v := range input {
		fmt.Println(v)
		fmt.Println(intToRoman(v))
	}
}