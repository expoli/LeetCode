package main

import "fmt"

func main() {
	//var num1 = []int{2, 4, 3}
	//var num2 = []int{5, 6, 4}

	//var num1 = []int{0}
	//var num2 = []int{0}

	//var num1 = []int{9, 9, 9, 9, 9, 9, 9}
	//var num2 = []int{9, 9, 9, 9}

	var num1 = []int{2, 4, 9}
	var num2 = []int{5, 6, 4, 9}

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
	var tail *ListNode
	var head *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		// 先进行数值初始化
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		// 求和计算
		sum := n1 + n2 + carry
		// 计算进位
		sum, carry = sum%10, sum/10
		// 头结点初始化操作
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			// 加入新结点
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	// 处理最高进位
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
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
