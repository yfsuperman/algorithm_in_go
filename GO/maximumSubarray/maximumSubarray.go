package maximumSubarray

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// Example:

// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Follow up:

// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    maxSum, sum := nums[0], 0
    
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        maxSum = findMax(maxSum, sum)
        sum = findMax(sum, 0)
    }
    
    return maxSum
}

func findMax(a, b int) int {
    if a > b {
        return a
    }
    return b
}