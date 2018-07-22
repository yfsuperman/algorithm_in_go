package longestSubstringWithoutRepeatingCharacters

// Given a string, find the length of the longest substring without repeating characters.

// Examples:

// Given "abcabcbb", the answer is "abc", which the length is 3.
// Given "bbbbb", the answer is "b", with the length of 1.
// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.


// This function can be improved by using slice instead of map to store the index of the characters
// Though the lookup cost is O(1) for map and O(n) for slice, the cost of hash functions for map 
// cannot be ignored when the input n is very small. (in this case, n <= 256) 
func lengthOfLongestSubstring(s string) int {
    
    if len(s) == 0 { return 0 }
    
    m := make(map[byte]int)
    res, start, end := 0, -1, 0
    
    for ; end < len(s); end++ {
        if index, ok := m[s[end]]; ok {
            if index > start { start = index }
        }
        m[s[end]] = end
        if end - start > res { res = end - start }
    }

    return res
}