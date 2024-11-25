package main

import "fmt"

func main() {
	obj := Constructor(5)
	param_1 := obj.InsertFront(7)
	param_2 := obj.InsertLast(0)
	param_3 := obj.GetFront()
	param_4 := obj.InsertLast(3)
	param_5 := obj.GetFront()
	param_6 := obj.InsertFront(9)
	param_7 := obj.GetRear()
	param_8 := obj.GetFront()
	param_9 := obj.GetFront()
	param_10 := obj.DeleteLast()
	param_11 := obj.GetRear()
	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
	fmt.Println(param_5)
	fmt.Println(param_6)
	fmt.Println(param_7)
	fmt.Println(param_8)
	fmt.Println(param_9)
	fmt.Println(param_10)
	fmt.Println(param_11)

}

type MyCircularDeque struct {
	size  int
	deque []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{size: k, deque: make([]int, 0, k)}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if !this.IsFull() {
		if this.IsEmpty() {
			this.deque = append(this.deque, value)
		} else {
			temp := make([]int, len(this.deque), cap(this.deque))
			copy(temp, this.deque)
			this.deque = make([]int, 0, this.size)
			this.deque = append(this.deque, value)
			this.deque = append(this.deque, temp...)
		}
		return true
	}
	return false
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if !this.IsFull() {
		this.deque = append(this.deque, value)
		return true
	}
	return false
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if len(this.deque) == 1 {
		this.deque = make([]int, 0, this.size)
	} else {
		this.deque = this.deque[1:]
	}
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.deque = this.deque[:len(this.deque)-1]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if !this.IsEmpty() {
		return this.deque[0]
	}
	return -1
}

func (this *MyCircularDeque) GetRear() int {
	if !this.IsEmpty() {
		return this.deque[len(this.deque)-1]
	}
	return -1
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.deque) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == len(this.deque)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
