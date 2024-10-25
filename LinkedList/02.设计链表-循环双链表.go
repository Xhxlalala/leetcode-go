package main

import "fmt"

type MyLinkedList struct {
	DummyHead *ListNode
}
type ListNode struct {
	Val  int
	Next *ListNode
	Pre  *ListNode
}

func Constructor() MyLinkedList {
	rear := &ListNode{ // 直接初始化结构体并返回指针
		Val:  -999,
		Next: nil,
		Pre:  nil,
	}
	rear.Next = rear
	rear.Pre = rear
	return MyLinkedList{rear}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.DummyHead.Next
	for cur != this.DummyHead && index > 0 {
		index--
		cur = cur.Next
	}
	if cur == this.DummyHead || index < 0 {
		return -1
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := new(ListNode)
	newNode.Val = val
	newNode.Next = this.DummyHead.Next
	newNode.Pre = this.DummyHead
	this.DummyHead.Next.Pre = newNode
	this.DummyHead.Next = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := new(ListNode)
	newNode.Val = val
	newNode.Next = this.DummyHead
	newNode.Pre = this.DummyHead.Pre
	this.DummyHead.Pre.Next = newNode
	this.DummyHead.Pre = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	head := this.DummyHead.Next
	for head != this.DummyHead && index > 0 {
		index--
		head = head.Next
	}
	if index > 0 {
		return
	}

	newNode := new(ListNode)
	newNode.Val = val
	newNode.Next = head
	newNode.Pre = head.Pre
	head.Pre.Next = newNode
	head.Pre = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	//链表为空
	if this.DummyHead.Next == this.DummyHead {
		return
	}
	cur := this.DummyHead.Next
	for cur.Next != this.DummyHead && index > 0 {
		index--
		cur = cur.Next
	}
	if index == 0 {
		cur.Pre.Next = cur.Next
		cur.Next.Pre = cur.Pre
	}
}

// 打印链表
func (list *MyLinkedList) printLinkedList() {
	// 从头节点开始遍历
	cur := list.DummyHead.Next
	// 当未达到哨兵节点时继续遍历
	for cur != list.DummyHead {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("END")
}

func main() {
	// 初始化链表
	list := Constructor()

	// 在头部添加元素
	list.AddAtHead(100)
	list.AddAtHead(200)
	fmt.Println("添加头部后：")
	list.printLinkedList()

	// 在尾部添加元素
	list.AddAtTail(300)
	list.AddAtTail(400)
	fmt.Println("添加尾部后：")
	list.printLinkedList()

	// 在指定位置添加元素
	list.AddAtIndex(2, 250)
	fmt.Println("在索引2处添加250后：")
	list.printLinkedList()

	// 删除指定位置的元素
	list.DeleteAtIndex(1)
	fmt.Println("删除索引1处的元素后：")
	list.printLinkedList()

	// 获取指定位置的元素
	val := list.Get(2)
	fmt.Printf("索引2处的元素值：%d\n", val)
}
