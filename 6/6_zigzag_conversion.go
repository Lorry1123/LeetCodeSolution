package main

import "fmt"
import "bytes"

func convert(s string, numRows int) string {
	f := make([][]byte, numRows)

	if numRows == 1 {
		return s
	}

	dir := 1
	row := 0
	for i := 0; i < len(s); i ++ {
		f[row] = append(f[row], s[i])
		
		if row == 0 {
			dir = 1
		}
		if row == numRows - 1 {
			dir = -1
		}
		
		row += dir
	}

	var ret bytes.Buffer
	for i := 0; i < numRows; i ++ {
		for j := 0; j < len(f[i]); j ++ {
			ret.WriteString(string(f[i][j]))
		}
	}
	return ret.String()
}

func main() {
	ret := convert("LEETCODEISHIRING", 3)
	fmt.Println(ret)

	ret = convert("LEETCODEISHIRING", 4)
	fmt.Println(ret)

	ret = convert("AB", 1)
	fmt.Println(ret)
}