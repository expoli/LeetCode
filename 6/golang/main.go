package golang

import (
	"fmt"
)

func convert(s string, numRows int) string {
	res := ""
	n := len(s)
	column := numRows
	mid := numRows - 2

	if n == 1 || numRows == 1 {
		return s
	}

	var temp2 string

	for i := 0; i < numRows; i++ {
		for j := 0; j < n; j++ {
			if j%(column+mid) == i || j%(column+mid) == (column+mid)-i {
				temp2 = fmt.Sprintf("%s%c", temp2, s[j])
			}
		}
	}
	res = temp2
	return res
}
