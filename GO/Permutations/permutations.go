package permutations

// Given a collection of distinct integers, return all possible permutations.

// Example:

// Input: [1,2,3]
// Output:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

func permute(nums []int) [][]int {
    res := make([][]int, 0)
    record := make([]bool, len(nums))
    helper(nums, []int{}, record, &res)
    return res
}

func helper(nums, tmp []int, record []bool, res *[][]int) {
    if len(tmp) == len(nums) {
        resTmp := make([]int, len(nums))
        copy(resTmp, tmp)
        *res = append(*res, resTmp)
        return
    }
    
    for i, val := range(nums) {
        if !record[i] {
            tmp = append(tmp, val)
            record[i] = true
            helper(nums, tmp, record, res)
            record[i] = false
            tmp = tmp[:len(tmp)-1]
        }
    }
}