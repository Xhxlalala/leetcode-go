package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func solution(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next
		tmp3 := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmp1
		cur.Next.Next.Next = tmp3
		cur = cur.Next.Next
	}
	return dummy.Next
}
