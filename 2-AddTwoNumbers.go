package main

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	val := l1.Val + l2.Val
	overflow := val / 10
	if overflow > 0 {
		return &ListNode{val % 10, AddTwoNumbers(AddTwoNumbers(&ListNode{overflow, nil}, l1.Next), l2.Next)}
	}
	return &ListNode{val, AddTwoNumbers(l1.Next, l2.Next)}
}