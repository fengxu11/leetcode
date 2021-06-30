package week01

import "fmt"

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
		fmt.Println("head.Next: ", head.Next)
		fmt.Println("last: ", last)

		// 改一边
		head.Next = last
		// last , head 向后移动一位
		last = head
		head = nextNode
	}

	return last
}
