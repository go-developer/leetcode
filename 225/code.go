// Package _25...
//
// Description : _25...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-10 5:40 下午
package main

func main() {

}

type MyStack struct {
	one []int
	two []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		one: make([]int, 0),
		two: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.one = append(this.one, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	this.fillTwo()
	data := this.two[len(this.two)-1]
	this.two = this.two[:len(this.two)-1]
	return data
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	this.fillTwo()
	return this.two[len(this.two)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.one) == 0 && len(this.two) == 0
}

// fillTwo 将one栈的数据填充进two栈
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:20 下午 2021/2/10
func (this *MyStack) fillTwo() {
	if len(this.one) == 0 {
		return
	}
	for i := 0; i < len(this.one); i++ {
		this.two = append(this.two, this.one[i])
	}
	this.one = make([]int, 0)
}
