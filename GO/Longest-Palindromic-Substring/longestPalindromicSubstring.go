package longestPalindromicSubstring

// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

// Example 1:

// Input: "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.
// Example 2:

// Input: "cbbd"
// Output: "bb"

func longestPalindrome(s string) string {
    if len(s) == 0 { return s }
    
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        len1 := findLength(s, i, i)
        len2 := findLength(s, i, i+1)
        maxLen := findMax(len1, len2)
        if maxLen > end - start {
            start = i - (maxLen - 1) / 2;
            end = i + maxLen / 2;
        }
    }
    
    return s[start:end+1]
    
}

func findMax(a int, b int) int {
    if a > b { return a }
    return b
}

func findLength(s string, l int, r int) int {
    for l >= 0 && r < len(s)  && s[l] == s[r] {
        l--
        r++
    }
    return r - l - 1
}