package combinationSumII

// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

// Each number in candidates may only be used once in the combination.

// Note:

// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:

// Input: candidates = [10,1,2,7,6,1,5], target = 8,
// A solution set is:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    return helper(candidates, target)
}

func helper(candidates []int, target int) [][]int {
    var res [][]int
    for i, val := range candidates {
        if i > 0 && candidates[i] == candidates[i-1] { continue }
        diff := target - val
        if diff > 0 {
            for _, arr := range helper(candidates[i+1:], diff) {
                tmp := append([]int{val}, arr...)
                res = append(res, tmp)
            }
        } else if diff == 0 {
            res = append(res, []int{val})
        }
    }
    
    return res
}