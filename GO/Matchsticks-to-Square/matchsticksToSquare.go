package matchsticksToSquare

// Remember the story of Little Match Girl? By now, you know exactly what matchsticks the little match girl has, please find out a way you can make one square by using up all those matchsticks. You should not break any stick, but you can link them up, and each matchstick must be used exactly one time.

// Your input will be several matchsticks the girl has, represented with their stick length. Your output will either be true or false, to represent whether you could make one square using all the matchsticks the little match girl has.

// Example 1:
// Input: [1,1,2,2,2]
// Output: true

// Explanation: You can form a square with length 2, one side of the square came two sticks with length 1.
// Example 2:
// Input: [3,3,3,3,4]
// Output: false

// Explanation: You cannot find a way to form a square with all the matchsticks.
// Note:
// The length sum of the given matchsticks is in the range of 0 to 10^9.
// The length of the given matchstick array will not exceed 15.

import (
	"sort"
)

func makesquare(nums []int) bool {
    if len(nums) < 4 { return false }
    sum, edge := 0, 0
    for _, i := range(nums) {
        sum += i
    }
    if (sum % 4 != 0) {
        return false
    } else {
        edge = sum/4
    }
    sort.Ints(nums)
    edges := make([]int, 4)
    return helper(nums, edges, len(nums)-1, edge)
}

func helper(nums, edges []int, index, edge int) bool {
    if index == 0 {
        if edges[0] == edge && edges[1] == edge && edges[2] == edge {
            return true
        }
        return false
    }
    
    for i := 0; i < 4; i++ {
        if edges[i] + nums[index] > edge { continue }
        edges[i] += nums[index]
        if helper(nums, edges, index-1, edge) { return true }
        edges[i] -= nums[index]
    }
    
    return false
}