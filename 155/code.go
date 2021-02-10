// Package _55...
//
// Description : _55...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-10 2:46 下午
package main

func main() {

}

// node 栈内每一个元素的结构
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:58 下午 2021/2/10
type node struct {
	current int // 当前值
	min     int // 最小值
}
type MinStack struct {
	stack []node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]node, 0),
	}
}

func (this *MinStack) Push(x int) {
	newNode := node{
		current: x,
		min:     x,
	}
	if this.IsEmpty() {
		this.stack = append(this.stack, newNode)
	} else {
		if this.stack[len(this.stack)-1].min < x {
			newNode.min = this.stack[len(this.stack)-1].min
		}
		this.stack = append(this.stack, newNode)
	}
}

func (this *MinStack) Pop() {
	if this.IsEmpty() {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if this.IsEmpty() {
		return 0
	}
	return this.stack[len(this.stack)-1].current
}

func (this *MinStack) GetMin() int {
	if this.IsEmpty() {
		return 0
	}
	return this.stack[len(this.stack)-1].min
}

func (this *MinStack) IsEmpty() bool {
	return len(this.stack) == 0
}
