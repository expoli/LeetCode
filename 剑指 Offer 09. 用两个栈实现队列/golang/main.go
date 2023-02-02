package golang

type CQueue struct {
	head int
	tail int
	data []int
}

type Stack struct {
	top  int
	data []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.data = append(this.data, value)
	this.head++
}

func (this *CQueue) DeleteHead() int {
	if this.head == this.tail {
		return -1
	}
	res := this.data[this.tail]
	this.tail++
	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
