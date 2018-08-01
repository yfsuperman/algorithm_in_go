package positionsOfLargeGroups

// In a string S of lowercase letters, these letters form consecutive groups of the same character.

// For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and "yy".

// Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.

// The final answer should be in lexicographic order.

 

// Example 1:

// Input: "abbxxxxzzy"
// Output: [[3,6]]
// Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.
// Example 2:

// Input: "abc"
// Output: []
// Explanation: We have "a","b" and "c" but no large group.

func largeGroupPositions(S string) [][]int {
    
    var res [][]int
    if len(S) == 0 { return res }
    
    prev, start, cnt := S[0], 0, 1
    for i := 1; i < len(S); i++ {
        if S[i] == prev {
            cnt++
        } else {
            if cnt >= 3 {
                res = append(res, []int{start, i-1})
            }
            prev, start, cnt = S[i], i, 1
        }
    }
    if cnt >= 3 { res = append(res, []int{start, len(S) - 1}) }
    return res
    
}