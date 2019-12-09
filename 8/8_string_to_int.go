package main

import "fmt"

import "math"

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func parseInt(c byte) int64 {
	return int64(c) - int64('0')
}

func myAtoi(str string) int {
	var ret int64 = 0
	var neg int64 = 1
	started := false

    for i := 0; i < len(str); i ++ {
		c := str[i]

		if !started {
			if c == ' ' {
				continue
			}
			if isNum(c) {
				ret = parseInt(c)
				started = true
				continue
			}
			if c == '-' {
				started = true
				neg = -1
				continue
			}
			if c == '+' {
				started = true
				neg = 1
				continue
			}

			return 0
		}

		if !isNum(c) {
			return int(ret)
		}
		tmp := ret * 10 + neg * parseInt(c)
		if tmp > math.MaxInt32 {
			return math.MaxInt32
		}
		if tmp < math.MinInt32 {
			return math.MinInt32
		}
		ret = tmp
	}
	
	if ret > math.MaxInt32 {
		return math.MaxInt32
	}
	if ret < math.MinInt32 {
		return math.MinInt32
	}
	return int(ret)
}

func main() {
	var ret int
	ret = myAtoi("42")
	fmt.Println(ret)

	ret = myAtoi("    -42")
	fmt.Println(ret)

	ret = myAtoi("4193 with words")
	fmt.Println(ret)

	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("+1"))
}