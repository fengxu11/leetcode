package week01

// 参考 leetcode 题解
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	// 维护一个前缀节点
	prehead := &ListNode{}
	result := prehead

	// 循环比较
	for l1 != nil && l2 != nil {
		// 让前缀节点指向 l1 或者 l2中较小的一个节点
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}

	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}
