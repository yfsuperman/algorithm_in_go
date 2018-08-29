package swapNodesInPairs

// Given a linked list, swap every two adjacent nodes and return its head.

// Example:

// Given 1->2->3->4, you should return the list as 2->1->4->3.

type ListNode struct {
    Val int
    Next *ListNode
}

/**
Solution 1: Loop
**/
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head}
    
    dummy := &ListNode {
    	Next: head,
    }
    prev, curt := dummy, head
    for curt != nil && curt.Next != nil {
        prev.Next = curt.Next
        curt.Next = curt.Next.Next
        prev.Next.Next = curt
        prev = curt
        curt = curt.Next
    }
    return dummy.Next
}

/**
Solution 2: recursion
**/
// func swapPairs(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil { return head }

//     next := head.Next
//     head.Next = swapPairs(next.Next)
//     next.Next = head
    
//     return next
// }