package validParentheses

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.

// Example 1:

// Input: "()"
// Output: true
// Example 2:

// Input: "()[]{}"
// Output: true

func isValid(s string) bool {
    var stack []int32
    
    for _, si := range s {
        switch si {
        case '(', '{', '[':
            stack = append(stack, si)
        case ')', '}', ']':
            lens := len(stack)
            if lens == 0 { return false }
            if (si == ')' && stack[lens-1] != '(') ||
            (si == '}' && stack[lens-1] != '{') ||
            (si == ']' && stack[lens-1] != '[') { return false }
            stack = stack[:lens-1]
        }
    }
    
    return len(stack) == 0 
}