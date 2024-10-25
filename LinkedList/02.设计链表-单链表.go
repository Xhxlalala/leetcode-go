package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	Size      int
	DummyHead *ListNode
}

func Constructor() MyLinkedList {
	newNode := &ListNode{ // 直接初始化结构体并返回指针
		Val:  -999,
		Next: nil,
	}
	return MyLinkedList{
		DummyHead: newNode,
		Size:      0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}
	cur := this.DummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := new(ListNode)
	newNode.Val = val
	newNode.Next = this.DummyHead.Next
	this.DummyHead.Next = newNode
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := new(ListNode)
	newNode.Val = val
	cur := this.DummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 { // 如果索引小于0，设置为0
		index = 0
	} else if index > this.Size { // 如果索引大于链表长度，直接返回
		return
	}

	newNode := new(ListNode)
	newNode.Val = val
	cur := this.DummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.DummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}

	this.Size--
}

// 打印链表
func (list *MyLinkedList) printLinkedList() {
	cur := list.DummyHead // 设置当前节点为虚拟头节点
	for cur.Next != nil { // 遍历链表
		fmt.Println(cur.Next.Val) // 打印节点值
		cur = cur.Next            // 切换到下一个节点
	}
}

func main() {
	list := Constructor()     // 初始化链表
	list.AddAtHead(100)       // 在头部添加元素
	list.AddAtTail(242)       // 在尾部添加元素
	list.AddAtTail(777)       // 在尾部添加元素
	list.AddAtIndex(1, 99999) // 在指定位置添加元素
	list.printLinkedList()    // 打印链表
}
