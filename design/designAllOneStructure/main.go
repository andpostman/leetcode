package main

import (
	"fmt"
	"math"
)

func main() {
	allOne := Constructor()
	// allOne.Inc("a")
	// allOne.Inc("b")
	// allOne.Inc("b")
	allOne.Inc("hello")
	allOne.Inc("hello")
	fmt.Println(allOne.GetMaxKey()) // return "hello"
	fmt.Println(allOne.GetMinKey()) // return "hello"
	allOne.Inc("leet")
	fmt.Println(allOne.GetMaxKey()) // return "hello"
	fmt.Println(allOne.GetMinKey()) // return "leet"

}

type Node struct {
	count int
	keys  map[string]bool
	prev  *Node
	next  *Node
}

type AllOne struct {
	keyToNode map[string]*Node
	head      *Node
	tail      *Node
}

func Constructor() AllOne {
	keyToNode := make(map[string]*Node)
	head := &Node{0, nil, nil, nil}
	tail := &Node{math.MaxInt32, nil, nil, nil}
	head.next = tail
	tail.prev = head
	return AllOne{keyToNode, head, tail}
}

func (this *AllOne) InsertNext(curr, next *Node) {
	currNext := curr.next
	curr.next = next
	currNext.prev = next
	next.prev = curr
	next.next = currNext
}

func (this *AllOne) InsertPrev(curr, prev *Node) {
	currPrev := curr.prev
	curr.prev = prev
	currPrev.next = prev
	prev.next = curr
	prev.prev = currPrev
}

func (this *AllOne) CleanUpNode(node *Node, key string) {
	delete(node.keys, key)
	if len(node.keys) == 0 {
		// Remove the node
		nodePrev, nodeNext := node.prev, node.next
		nodePrev.next = nodeNext
		nodeNext.prev = nodePrev
	}
}

func (this *AllOne) Inc(key string) {
	// If new key
	if this.keyToNode[key] == nil {
		// If Node with count 1 does not exist, create it
		if this.head.next.count != 1 {
			node := &Node{1, make(map[string]bool), nil, nil}
			// Insert key reference
			node.keys[key] = true
			// Attach new Node to head
			this.InsertNext(this.head, node)
			// Insert key to Node reference
			this.keyToNode[key] = node
		} else {
			node := this.head.next
			node.keys[key] = true
			this.keyToNode[key] = node
		}
	} else { // Existing key
		node := this.keyToNode[key]
		// Create node with count + 1
		if node.next.count != node.count+1 {
			nodePlusOne := &Node{node.count + 1, make(map[string]bool), nil, nil}
			nodePlusOne.keys[key] = true
			this.InsertNext(node, nodePlusOne)

			// Remove key from old node
			this.CleanUpNode(node, key)

			this.keyToNode[key] = nodePlusOne
		} else {
			nodePlusOne := node.next
			nodePlusOne.keys[key] = true
			// Remove key from old node
			this.CleanUpNode(node, key)

			this.keyToNode[key] = nodePlusOne
		}
	}
}

func (this *AllOne) Dec(key string) {
	node := this.keyToNode[key]
	if node == nil {
		return
	}
	if node.count == 1 {
		this.CleanUpNode(node, key)
		delete(this.keyToNode, key)
	}
	if node.count > 1 {
		if node.prev.count != node.count-1 {
			nodeMinusOne := &Node{node.count - 1, make(map[string]bool), nil, nil}
			nodeMinusOne.keys[key] = true
			this.InsertPrev(node, nodeMinusOne)

			// Remove key from old node
			this.CleanUpNode(node, key)

			this.keyToNode[key] = nodeMinusOne
		} else {
			nodeMinusOne := node.prev
			nodeMinusOne.keys[key] = true
			// Remove key from old node
			this.CleanUpNode(node, key)

			this.keyToNode[key] = nodeMinusOne
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.tail.prev == this.head {
		return ""
	}
	var res string
	for key := range this.tail.prev.keys {
		res = key
		break
	}
	return res
}

func (this *AllOne) GetMinKey() string {
	if this.head.next == this.tail {
		return ""
	}
	var res string
	for key := range this.head.next.keys {
		res = key
		break
	}
	return res
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
