package main

import "fmt"

type MinStack struct {
	top  int
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if this.top == 0 {
		this.min = append(this.min, x)
	} else {
		if x < this.min[this.top-1] {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, this.min[this.top-1])
		}
	}
	this.data = append(this.data, x)
	this.top++
}

func (this *MinStack) Pop() {
	this.top--
	this.min = this.min[:this.top]
	this.data = this.data[:this.top]
}

func (this *MinStack) Top() int {
	return this.data[this.top-1]
}

func (this *MinStack) Min() int {
	return this.min[this.top-1]
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
