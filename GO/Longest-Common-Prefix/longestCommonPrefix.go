package longestCommonPrefix

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: ["flower","flow","flight"]
// Output: "fl"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 { return "" }
    maxLen := findShortestLength(strs)
    
    for i := 0; i < maxLen; i++ {
        ithChar := strs[0][i]
        for _, str := range strs {
            if ithChar != str[i] {
                return strs[0][:i]
            }
        }
    }
    
    return strs[0][:maxLen]
}

func findShortestLength(strs []string) int {
    if len(strs) == 0 { return 0 }
    
    res := len(strs[0])
    for _, str := range strs {
        if len(str) < res {
            res = len(str)
        }
    }
    
    return res
}