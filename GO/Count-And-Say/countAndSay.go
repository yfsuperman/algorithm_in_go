package countAndSay

import (
	"strconv"
)

// The count-and-say sequence is the sequence of integers with the first five terms as following:

// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221
// 1 is read off as "one 1" or 11.
// 11 is read off as "two 1s" or 21.
// 21 is read off as "one 2, then one 1" or 1211.
// Given an integer n, generate the nth term of the count-and-say sequence.

func countAndSay(n int) string {
    if n < 1 { return "" }
    
    result := "1"
    for i := 1; i < n; i++ {
        result = findNext(result)
    }
    
    return result
}

func findNext(s string) string {
    var res string
    curt, cnt := s[0], 1
    for i := 1; i < len(s); i++ {
        if s[i] == curt { 
            cnt++ 
        } else {  
            res = res + strconv.Itoa(cnt) + string(curt)
            curt, cnt = s[i], 1
        }
    }
    res = res + strconv.Itoa(cnt) + string(curt)
    return res
}