package groupAnagrams

// Given an array of strings, group anagrams together.

// Example:

// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
// Note:

// All inputs will be in lowercase.
// The order of your output does not matter.

import (
    "sort"
)

func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    
    for _, str := range strs {
        chars := []rune(str)
        sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
        key := string(chars)
        m[key] = append(m[key], str)
    }
    
    var res [][]string
    for _, val := range m {
        res = append(res, val)
    }
    
    return res
    
}