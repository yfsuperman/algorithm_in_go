package threeSumClosest

// Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

// Example:

// Given array nums = [-1, 2, 1, -4], and target = 1.

// The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

import (
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
    res, sum := math.MaxInt32, math.MaxInt32
    sort.Ints(nums)
    
    for i := 0; i < len(nums) -2; i++ {
        if i > 0 && nums[i] == nums[i-1] {continue}
        low, high := i+1, len(nums)-1
        
        for low < high {
            tmpSum := nums[i] + nums[low] + nums[high]
            diff := target - tmpSum
            if getAbs(diff) < res {
                res, sum = getAbs(diff), tmpSum
            }
            if diff == 0 {
                return target
            } else if diff > 0 {
                low ++
            } else {
                high --
            }
        }
    }
    
    return sum
}

func getAbs(a int) int {
    if a < 0 { return -a }
    
    return a
}