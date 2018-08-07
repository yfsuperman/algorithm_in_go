package threeSum

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

// Note:
// The solution set must not contain duplicate triplets.

// Example:
// Given array nums = [-1, 0, 1, 2, -1, -4],
// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

import (
    "sort"
)
func threeSum(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    
    for i := 0; i < len(nums) - 2; i++ {
        if i > 0 && nums[i] == nums[i-1] { continue }
        
        low, high := i+1, len(nums)-1
        for low < high {
            lowVal, highVal := nums[low], nums[high]
            if lowVal + highVal + nums[i] == 0 {
                res = append(res, []int{nums[i], lowVal, highVal})
                for low < high && nums[low] == lowVal {
                    low++
                }
                for low < high && nums[high] == highVal {
                    high--
                }
            } else if lowVal + highVal + nums[i] > 0 {
                high--
            } else {
                low++
            }
        }
    }
    
    return res
    
}