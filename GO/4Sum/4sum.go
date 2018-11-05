package fourSum

// Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

// Note:

// The solution set must not contain duplicate quadruplets.

// Example:

// Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

// A solution set is:
// [
//   [-1,  0, 0, 1],
//   [-2, -1, 1, 2],
//   [-2,  0, 0, 2]
// ]

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for last := len(nums) - 1; last > first+2; last-- {
			if last < len(nums)-1 && nums[last] == nums[last+1] {
				continue
			}
			twoSum(nums[first+1:last], target-nums[first]-nums[last], nums[first], nums[last], &res)
		}
	}

	return res
}

func twoSum(nums []int, target int, first int, last int, res *[][]int) {
	second, third := 0, len(nums)-1
	for second < third {
		if second > 0 && nums[second] == nums[second-1] {
			second++
			continue
		}
		if third < len(nums)-1 && nums[third] == nums[third+1] {
			third--
			continue
		}
		if nums[second]+nums[third] == target {
			*res = append(*res, []int{first, nums[second], nums[third], last})
			second++
			third--
		} else if nums[second]+nums[third] > target {
			third--
		} else {
			second++
		}
	}

}
