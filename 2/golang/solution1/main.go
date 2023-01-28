package main

import "fmt"

func main() {
	//var num1 = []int{2, 4, 3}
	//var num2 = []int{5, 6, 4}

	//var num1 = []int{0}
	//var num2 = []int{0}

	var num1 = []int{9, 9, 9, 9, 9, 9, 9}
	var num2 = []int{9, 9, 9, 9}
	l1 := createANumber(num1)
	l2 := createANumber(num2)

	printList(l1)
	fmt.Println("----------------------")
	printList(l2)
	fmt.Println("----------------------")
	sum := addTwoNumbers(l1, l2)
	printList(sum)
}

// ListNode
/*
单向链表定义
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(header *ListNode) {
	var temp = header
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var temp []int
	var flag = 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val
		if flag == 1 {
			sum += 1
			flag = 0
		}
		if sum > 9 {
			flag = 1
			temp = append(temp, sum%10)
		} else {
			temp = append(temp, sum)
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	// 进行后续数据结点的处理
	var lastList *ListNode
	if l1 != nil {
		lastList = l1
	} else {
		lastList = l2
	}
	// 求和计算
	for lastList != nil {
		sum := lastList.Val
		if flag == 1 {
			sum += 1
			flag = 0
		}
		if sum > 9 {
			flag = 1
			temp = append(temp, sum%10)
		} else {
			temp = append(temp, sum)
		}
		lastList = lastList.Next
	}
	// 最高位有进位的情况
	if flag == 1 {
		temp = append(temp, 1)
	}
	// 根据切片创建链表，这个是否能够优化掉呢？（比如一边计算一边构建链表）
	number := createANumber(temp)
	return number
}

/*
根据输入的整数切片进行数字链表的构建，头结点进行特殊处理
*/
func createANumber(nums []int) *ListNode {
	var header ListNode
	var prev *ListNode
	for i, num := range nums {
		// 初始化头结点
		if i == 0 {
			header.Val = num
			header.Next = nil
			prev = &header
		} else {
			// 向链表中添加其它的数据结点
			var temp = ListNode{
				Val:  num,
				Next: nil,
			}
			prev.Next = &temp
			prev = &temp
		}
	}
	return &header
}
