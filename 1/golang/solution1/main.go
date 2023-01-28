package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

/*
*
暴力解法：即两个 for 循环
- 分析
  - 时间复杂度
  - O(n^2)
  - 空间复杂度
  - O(1)
*/
func twoSum(nums []int, target int) []int {
	for i, temp := range nums {
		for j := i + 1; j < len(nums); j++ {
			if temp+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
