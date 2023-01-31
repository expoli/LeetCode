package main

import "fmt"

func main() {
	test1 := [][]int{{5, 7, 0}, {0, 3, 1}, {0, 5, 0}}
	test2 := [][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}}

	// true
	fmt.Println(checkXMatrix(test2))
	// false
	fmt.Println(checkXMatrix(test1))
}

func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	m := len(grid[0])
	// 如果不是方阵直接返回
	if n == 0 || m != n {
		return false
	}
	for i, row := range grid {
		// 对角线元素判断
		if row[i] == 0 || row[n-1-i] == 0 {
			return false
		}
		for j := 0; j < n; j++ {
			if j == i || j == n-1-i {
				continue
			}
			if row[j] != 0 {
				return false
			}
		}
	}
	return true
}
