package addTwoNumbers

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example:

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

type ListNode struct {
    Val int
    Next *ListNode
}

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := ListNode{0, nil}
    var dummy *ListNode = &res
    count := 0
    tmp := 0
    for (l1 != nil || l2 != nil || count != 0) {
        var a, b int
        if l1 != nil {
            a = l1.Val
        } else {
            a = 0 
        }
        if l2 != nil {
            b = l2.Val
        } else {
            b = 0 
        }
        tmp = (a + b + count) % 10
        count = (a + b + count) / 10
        p := ListNode{tmp, nil}
        dummy.Next = &p
        dummy = dummy.Next
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }
    
    return res.Next
}