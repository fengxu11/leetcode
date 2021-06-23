package week01

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {

	var last *ListNode = nil

	for head != nil {
		// 暂存
		nextNode := head.Next
		// 改一边
		head.Next = last
		// last , head 向后移动一位
		last = head
		head = nextNode
	}

	return last
}
