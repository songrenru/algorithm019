/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
type MyCircularDeque struct {
	capacity, count int
	head, tail *DLinkedNode
}


type DLinkedNode struct {
	val int
	prev, next *DLinkedNode
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	deque := MyCircularDeque{
		capacity: k,
		head: &DLinkedNode{val:0},
		tail: &DLinkedNode{val:0},
	}
	deque.head.next = deque.tail
	deque.tail.prev = deque.head
	return deque
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &DLinkedNode{val:value}
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	this.count++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &DLinkedNode{val:value}
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev.next = node
	this.tail.prev = node
	this.count++
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	temp := this.head.next.next
	this.head.next = temp
	temp.prev = this.head
	this.count--
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	temp := this.tail.prev.prev
	this.tail.prev = temp
	temp.next = this.tail
	this.count--
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.next.val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.prev.val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.count == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.count == this.capacity
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
// @lc code=end

