package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	head, _ = reverseListExp(head)
	return head
}

func reverseListExp(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	current := &ListNode{head.Val, nil}

	if head.Next == nil {
		return current, current
	}

	head, end := reverseListExp(head.Next)
	// add current number to the end of list
	end.Next = current

	// return the head and end of the list
	return head, end.Next
}
