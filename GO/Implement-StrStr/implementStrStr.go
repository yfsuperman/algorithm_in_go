package implementStrStr

// Implement strStr().

// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

// Example 1:

// Input: haystack = "hello", needle = "ll"
// Output: 2
// Example 2:

// Input: haystack = "aaaaa", needle = "bba"
// Output: -1

func strStr(haystack string, needle string) int {
    h, n := len(haystack), len(needle)
    for i := 0; i < h - n + 1; i++ {
        if haystack[i:i+n] == needle {
            return i
        }
    } 
    return -1
}