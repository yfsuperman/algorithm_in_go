package targetSum

// You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols + and -. For each integer, you should choose one from + and - as its new symbol.

// Find out how many ways to assign symbols to make sum of integers equal to target S.

// Example 1:
// Input: nums is [1, 1, 1, 1, 1], S is 3. 
// Output: 5
// Explanation: 

// -1+1+1+1+1 = 3
// +1-1+1+1+1 = 3
// +1+1-1+1+1 = 3
// +1+1+1-1+1 = 3
// +1+1+1+1-1 = 3

// There are 5 ways to assign symbols to make the sum of nums be target 3.


func findTargetSumWays(nums []int, S int) int {
    cnt := 0
    helper(nums, S, 0, &cnt)
    return cnt
    
}

func helper(nums []int, S, curt int, cnt *int) {
    if len(nums) == 0 {
        if curt == S { *cnt ++ }
        return
    }
    
    helper(nums[1:], S, curt+nums[0], cnt)
    helper(nums[1:], S, curt-nums[0], cnt)
}