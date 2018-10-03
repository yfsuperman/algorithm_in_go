package wordLadder

// Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

// Only one letter can be changed at a time.
// Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
// Note:

// Return 0 if there is no such transformation sequence.
// All words have the same length.
// All words contain only lowercase alphabetic characters.
// You may assume no duplicates in the word list.
// You may assume beginWord and endWord are non-empty and are not the same.
// Example 1:

// Input:
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]

// Output: 5

// Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
// return its length 5.

func ladderLength(beginWord string, endWord string, wordList []string) int {
    if beginWord == endWord { return 0 }
    
    dict := make(map[string]bool, len(wordList))
    queue := make([]string, 0)
    for _, i := range wordList {
        if i != beginWord {
            dict[i] = true
        }
    }
    
    queue = append(queue, beginWord)
    
    cnt := 1
    for len(queue) > 0 {
        queueLen := len(queue)
        for q := 0; q < queueLen; q++ {
            curt := queue[0]
            queue = queue[1:]
            bytes := []byte(curt)
            for i := 0; i < len(bytes); i++ {
                oldChar := bytes[i]
                for j := 0; j < 26; j++ {
                    newChar := 'a' + byte(j)
                    if newChar == oldChar {
                        continue
                    }
                    bytes[i] = newChar
                    // newWord := string(bytes) // This helps reduces 72ms if we do not create new variable
                    if val, ok := dict[string(bytes)]; ok {
                        if string(bytes) == endWord { return cnt+1 }
                        if val == true {
                            queue = append(queue, string(bytes))
                            dict[string(bytes)] = false
                        }
                    }
                }
                bytes[i] = oldChar
            }
        }
        cnt++
    }
    
    return 0
    
}