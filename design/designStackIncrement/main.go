package main

import (
	"fmt"
)

func main() {
	stk := Constructor(3) // Stack is Empty []
	stk.Push(1)           // stack becomes [1]
	fmt.Println(stk.stack)
	stk.Push(2) // stack becomes [1, 2]
	fmt.Println(stk.stack)
	fmt.Println(stk.Pop()) // return 2 --> Return top of the stack 2, stack becomes [1]
	fmt.Println(stk.stack)
	stk.Push(2) // stack becomes [1, 2]
	fmt.Println(stk.stack)
	stk.Push(3) // stack becomes [1, 2, 3]
	fmt.Println(stk.stack)
	stk.Push(4) // stack still [1, 2, 3], Do not add another elements as size is 4
	fmt.Println(stk.stack)
	stk.Increment(5, 100) // stack becomes [101, 102, 103]
	fmt.Println(stk.stack)
	stk.Increment(2, 100) // stack becomes [201, 202, 103]
	fmt.Println(stk.stack)
	fmt.Println(stk.Pop()) // return 103 --> Return top of the stack 103, stack becomes [201, 202]
	fmt.Println(stk.stack)
	fmt.Println(stk.Pop()) // return 202 --> Return top of the stack 202, stack becomes [201]
	fmt.Println(stk.stack)
	fmt.Println(stk.Pop()) // return 201 --> Return top of the stack 201, stack becomes []
	fmt.Println(stk.stack)
	fmt.Println(stk.Pop())
	fmt.Println(stk.stack)
}

type CustomStack struct {
	length int
	stack  []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		length: maxSize,
		stack:  make([]int, 0, maxSize),
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.stack) != this.length {
		this.stack = append(this.stack, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.stack) == 0 {
		return -1
	}
	old := this.stack
	n := len(old)
	item := old[n-1]
	this.stack = old[0 : n-1]
	return item
}

func (this *CustomStack) Increment(k int, val int) {
	var size int
	if k > len(this.stack) {
		size = len(this.stack)
	} else {
		size = k
	}
	for i := 0; i < size; i++ {
		this.stack[i] += val
	}
}

/**
* Your CustomStack object will be instantiated and called as such:
* obj := Constructor(maxSize);
* obj.Push(x);
* param_2 := obj.Pop();
* obj.Increment(k,val);
 */
