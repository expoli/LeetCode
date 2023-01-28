package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

/*
*
使用哈希表来进行目标元素的查找
如果哈希表中不存在则进行插入操作进行缓存

第一次：	i -> 0, x -> 2	target-x -> 7	hashTable[7] -> nil	hashTable[2] -> 0
第二次：	i -> 1, x -> 7	target-x -> 2	hashTable[2] -> 0
输出：[0 1]

- 代码分析
  - 时间复杂度
  - O(n)
  - 空间复杂度
  - O(n)
*/
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
