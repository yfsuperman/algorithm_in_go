package combinationSum

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

// The same repeated number may be chosen from candidates unlimited number of times.

// Note:

// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:

// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]

func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
    for i, val := range candidates {
        diff := target - val
        if diff > 0 {
            for _, arr := range combinationSum(candidates[i:], diff) {
				// Adding dots after second slice can let you pass multiple arguments to a variadic function from a slice
                tmp := append([]int{val}, arr...)
                res = append(res, tmp)
            }          
        } else if diff == 0 {
            res = append(res, []int{val})
        }
    }
    
    return res
}
