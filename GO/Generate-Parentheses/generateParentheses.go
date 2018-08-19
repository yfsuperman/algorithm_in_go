package generateParentheses

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

// For example, given n = 3, a solution set is:

// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

/*
Solution 1
*/

func generateParenthesis(n int) []string {
    res := []string{}
    helper("", &res, n, n)
    return res
}

func helper (curt string, res *[]string, l, r int) {
    if l < 0 || r < 0 { return }
    
    if l == 0 && r == 0 {
        *res = append(*res, curt)
        return
    }
    
    if l > 0 {
        helper(curt + "(", res, l-1, r)
    }
    if r > l {
        helper(curt + ")", res, l, r-1)
    }
}

/*
Solution 2
*/
// func generateParenthesis(n int) []string {
//     return helper("", n, n)
// }

// func helper (curt string, l, r int) []string {
//     if l < 0 || r < 0 || l > r { return nil }
    
//     if l == 0 && r == 0 {
//         return []string{curt}
//     }
    
//     if l == r {
//         return helper(curt + "(", l-1, r)
//     }
    
//     res := []string{}
//     if l > 0 {
//         res = append(res, helper(curt + "(", l-1, r)...)
//     }
//     if r > 0 {
//         res = append(res, helper(curt + ")", l, r-1)...)
//     }
//     return res
// }