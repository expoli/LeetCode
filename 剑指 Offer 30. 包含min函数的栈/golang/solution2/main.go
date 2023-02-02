package main

import (
	"fmt"
	"math"
)

type MinStack struct {
}
type Node struct {
	min  int
	val  int
	next *Node
}

var head *Node = nil

/** initialize your data structure here. */
func Constructor() MinStack {
	head = nil
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if head == nil {
		head = &Node{
			min:  x,
			val:  x,
			next: nil,
		}
	} else {
		head = &Node{
			min:  int(math.Min(float64(x), float64(head.min))),
			val:  x,
			next: head,
		}
	}
}

func (this *MinStack) Pop() {
	head = head.next
}

func (this *MinStack) Top() int {
	return head.val
}

func (this *MinStack) Min() int {
	return head.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.Min())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.Min())
}
