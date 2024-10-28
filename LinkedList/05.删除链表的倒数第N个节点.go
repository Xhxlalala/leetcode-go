package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func solve(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{0, head}
	slow, fast := dummyHead, dummyHead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
