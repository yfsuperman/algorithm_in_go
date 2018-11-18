package substringWithConcatenationOfAllWords

// You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

// Example 1:

// Input:
//   s = "barfoothefoobarman",
//   words = ["foo","bar"]
// Output: [0,9]
// Explanation: Substrings starting at index 0 and 9 are "barfoor" and "foobar" respectively.
// The output order does not matter, returning [9,0] is fine too.
// Example 2:

// Input:
//   s = "wordgoodstudentgoodword",
//   words = ["word","student"]
// Output: []

func findSubstring(s string, words []string) []int {
    res := make([]int, 0)
    if len(words) == 0 {
        return res
    }
    wordMap := make(map[string]int, 0)
    n, m := len(words), len(words[0])
    for _, word := range(words) {
        wordMap[word] = wordMap[word] + 1
    }
    
    totalLen := n * m
    for i := 0; i <= len(s) - totalLen; i++ {
        if checkRes(s[i:i+totalLen], wordMap, m) {
            res = append(res, i)
        }
    }
    
    return res
    
}

func checkRes(s string, wordMap map[string]int, m int) bool {
    tmpMap := make(map[string]int, 0)
    for i := 0; i <= len(s) - m; i += m {
        tmp := s[i:i+m]
        if val, ok := tmpMap[tmp]; ok {
            if val + 1 > wordMap[tmp] {
                return false
            }
            tmpMap[tmp] = val + 1
        } else {
            if wordMap[tmp] == 0 {
                return false
            }
            tmpMap[tmp] = 1
        }
    }
    
    return true
}