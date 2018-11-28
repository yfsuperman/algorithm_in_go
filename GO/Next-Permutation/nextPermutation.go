package nextPermutation

// Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

// If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

// The replacement must be in-place and use only constant extra memory.

// Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1

func nextPermutation(nums []int)  {
    n := len(nums)
    if n == 1 {
        return
    }
    for i := n - 2; i >= 0; i-- {
        for j := n-1; j > i; j-- {
            if nums[j] > nums[i] {
                nums[i], nums[j] = nums[j], nums[i]
                sortArray(nums, i+1, n-1)
                return
            }
        }
    }
    
    sortArray(nums, 0, n-1)
}

func sortArray(nums []int, s, e int) {
    for s < e {
        nums[s], nums[e] = nums[e], nums[s]
        s ++
        e --
    }
}