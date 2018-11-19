package reverseNodesInKGroup

// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

// Example:

// Given this linked list: 1->2->3->4->5

// For k = 2, you should return: 2->1->4->3->5

// For k = 3, you should return: 3->2->1->4->5

// Note:

// Only constant extra memory is allowed.
// You may not alter the values in the list's nodes, only nodes itself may be changed.

/**
 * Definition for singly-linked list.
*/
type ListNode struct {
    Val int
    Next *ListNode
}

 func reverseKGroup(head *ListNode, k int) *ListNode {
    if !checkNeedReverse(head, k) {
        return head
    }
    res := ListNode{
        Val: 0,
        Next: head,
    }
    curt, next := head, head.Next
    
    cnt := 1
    for cnt < k && next != nil {
        post := next.Next
        next.Next = curt
        curt, next = next, post
        cnt ++
    }
    head.Next = reverseKGroup(next, k)
    res.Next = curt
    
    return res.Next
}

func checkNeedReverse(head *ListNode, k int) bool {
    cnt := 0
    for head != nil {
        cnt ++
        head = head.Next
    }
    if cnt < k {
        return false
    }
    return true
}