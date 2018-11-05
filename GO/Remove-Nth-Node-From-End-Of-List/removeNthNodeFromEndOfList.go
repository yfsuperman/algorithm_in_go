package removeNthNodeFromEndOfList

// Given a linked list, remove the n-th node from the end of list and return its head.

// Example:

// Given linked list: 1->2->3->4->5, and n = 2.

// After removing the second node from the end, the linked list becomes 1->2->3->5.
// Note:

// Given n will always be valid.

/*
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	res := dummy
	listLen := getListLen(head)
	for i := 0; i < listLen-n; i++ {
		dummy = dummy.Next
	}
	dummy.Next = dummy.Next.Next

	return res.Next

}

func getListLen(head *ListNode) int {
	cnt := 0
	for head != nil {
		cnt++
		head = head.Next
	}

	return cnt
}

/*
 * Solution 2: one loop
 */
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	if nil == head {
// 		return nil
// 	}
// 	ret := &ListNode{0, nil}
// 	ret.Next = head
// 	pre, cur := ret, ret

// 	for n > 0 && nil != pre.Next {
// 		pre = pre.Next
// 		n--
// 	}
// 	for nil != pre.Next {
// 		pre = pre.Next
// 		cur = cur.Next
// 	}
// 	pre = cur.Next
// 	cur.Next = cur.Next.Next
// 	return ret.Next
// }
