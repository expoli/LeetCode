package main

import "fmt"

type MinStack struct {
	top    int
	minTop int
	data   map[int]int
	min    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var s1 MinStack
	s1.data = make(map[int]int)
	return s1
}

func (this *MinStack) Push(x int) {
	if this.top == 0 {
		this.min = append(this.min, x)
		this.minTop++
	} else {
		if x <= this.min[this.minTop-1] {
			this.min = append(this.min, x)
			this.minTop++
		}
	}
	this.data[this.top] = x
	this.top++
}

func (this *MinStack) Pop() {
	this.top--
	temp := this.data[this.top]
	if temp == this.min[this.minTop-1] {
		this.minTop--
		this.min = this.min[:this.minTop]
	}
	delete(this.data, this.top)
}

func (this *MinStack) Top() int {
	return this.data[this.top-1]
}

func (this *MinStack) Min() int {
	return this.min[this.minTop-1]
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
