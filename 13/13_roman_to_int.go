package main

import "fmt"

import "strings"

func romanToInt(s string) int {
	m := map[string]int{
		"M": 1000,
		"CM": 900,
		"D": 500,
		"CD": 400,
		"C": 100,
		"XC": 90,
		"L": 50,
		"XL": 40,
		"X": 10,
		"IX": 9,
		"V": 5,
		"IV": 4,
		"I": 1,
	}

	iter := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ret := 0
	for _, v := range iter {
		for strings.HasPrefix(s, v) {
			ret += m[v]
			s = s[len(v):]
		}
		if len(s) <= 0 {
			break
		}
	}

	return ret
}

func main() {
	input := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	for _, v := range input {
		fmt.Println(v)
		fmt.Println(romanToInt(v))
	}
}