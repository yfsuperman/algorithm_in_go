package longestValidParentheses

// Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

// Example 1:

// Input: "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()"
// Example 2:

// Input: ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()"

type parentheses struct {
    index int
    char byte  
}

func longestValidParentheses(s string) int {
    stack := make([]parentheses, 0)
    res := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            stack = append(stack, parentheses{
                index: i, 
                char: s[i],
            })
        } else {
            if (len(stack) > 0) && (stack[len(stack)-1].char == '(') {
                cnt := 0
                stack = stack[:len(stack)-1]
                if len(stack) == 0 {
                    cnt = i + 1
                } else {
                    cnt = i - stack[len(stack)-1].index
                }
                res = findMax(res, cnt)
            } else {
                stack = append(stack, parentheses{
                    index: i, 
                    char: s[i],
                })
            }
        }
    }
    
    return res
}

func findMax(a, b int) int {
    if a > b {
        return a
    }
    return b
}