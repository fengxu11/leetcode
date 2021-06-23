package week01

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {

	fmt.Println("创建后: ")
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}

	cur := head
	for cur != nil {
		fmt.Println("value: ", cur.Val)
		cur = cur.Next
	}

	fmt.Println("反转后: ")
	head = reverseLinkedList(head)

	cur = head
	for cur != nil {
		fmt.Println("value: ", cur.Val)
		cur = cur.Next
	}

}
