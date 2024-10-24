package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 不带虚拟头节点
func removeElements1(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 带虚拟头结点
func removeElements2(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
