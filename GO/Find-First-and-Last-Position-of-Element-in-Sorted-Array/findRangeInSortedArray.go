package findRangeInSortedArray

// Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

// Your algorithm's runtime complexity must be in the order of O(log n).

// If the target is not found in the array, return [-1, -1].

// Example 1:

// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
// Example 2:

// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]

func searchRange(nums []int, target int) []int {
    return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
    res := -1
    if len(nums) == 0 {
        return res
    }
    start, end := 0, len(nums) - 1
    for start <= end {
        mid := start + (end - start) / 2
        if nums[mid] >= target {
            end = mid - 1
        } else {
            start = mid + 1
        }
    }
    if end >= 0 && nums[end] == target {
        res = end
    }
    if start < len(nums) && nums[start] == target {
        res = start
    }
    
    return res
}

func findRight(nums []int, target int) int {
    res := -1
    if len(nums) == 0 {
        return res
    }
    start, end := 0, len(nums) - 1
    for start <= end {
        mid := start + (end - start) / 2
        if nums[mid] > target {
            end = mid - 1
        } else {
            start = mid + 1
        }
    }
    if start < len(nums) && nums[start] == target {
        res = start
    }
    if end >= 0 && nums[end] == target {
        res = end
    }
    
    return res
}