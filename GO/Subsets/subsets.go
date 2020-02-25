package Subsets

// Given a set of distinct integers, nums, return all possible subsets (the power set).

// Note: The solution set must not contain duplicate subsets.

// Example:

// Input: nums = [1,2,3]
// Output:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

func subsets(nums []int) [][]int {
    res := [][]int{}
    if len(nums) == 0 {
		return res
	}
    
    backtracking(nums, []int{}, &res)
    return res
}

func backtracking(nums, curt []int, res *[][]int) {
    *res = append(*res, curt)

	for i := 0; i < len(nums); i++ {
		next := append([]int{}, curt...)
		backtracking(nums[i+1:], append(next, nums[i]), res)
	}
}