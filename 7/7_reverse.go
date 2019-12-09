package main

import "fmt"
import "math"

func reverse(x int) int {
	var tmp, ret int64
	if x > 0 {
		tmp = int64(x)
	} else {
		tmp = int64(-x)
	}

	ret = 0
	for tmp > 0 {
		ret = ret * 10 + tmp % 10
		tmp = tmp / 10
	}

	if x < 0 {
		ret *= -1
	}

	if ret > math.MaxInt32 {
		return 0
	} else if ret < math.MinInt32 {
		return 0
	} else {
		return int(ret)
	}
}

func main() {
	var ret int
	ret = reverse(123)
	fmt.Println(ret)

	ret = reverse(-123)
	fmt.Println(ret)

	ret = reverse(120)
	fmt.Println(ret)
}