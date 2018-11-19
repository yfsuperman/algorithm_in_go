package mergeKSortedLists

// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

// Example:

// Input:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// Output: 1->1->2->3->4->4->5->6

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
    listLen := len(lists)
    if listLen == 0 {
        return nil
    }
    if listLen == 1 {
        return lists[0]
    }
    
    if listLen == 2 {
        return mergeTwoLists(lists[0], lists[1])
    }
    
    return mergeKLists([]*ListNode{mergeKLists(lists[listLen/2:]), mergeKLists(lists[:listLen/2])})
}

func mergeTwoLists(a, b *ListNode) *ListNode {
    curt := &ListNode{
        Val: 0,
        Next: nil,
    }
    res := curt
    for a != nil && b != nil {
        if a.Val < b.Val {
            curt.Next, a = a, a.Next
        } else {
            curt.Next, b = b, b.Next
        }
        curt = curt.Next
    }
    
    if a == nil {
        curt.Next = b
    }
    if b == nil {
        curt.Next = a
    }
    return res.Next
}