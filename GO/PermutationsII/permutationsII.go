package permutationsII

// Given a collection of numbers that might contain duplicates, return all possible unique permutations.

// Example:

// Input: [1,1,2]
// Output:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	record := make([]bool, n)
	vector := make([]int, 0)
	helper(nums, vector, record, &res)
	return res
}

/**
** Solution1:
**/
// func helper(nums, vector []int, record []bool, res *[][]int) {
//     curtLen, n := len(vector), len(nums)
//     if curtLen == len(nums) {
//         tmp := make([]int, n)
//         copy(tmp, vector)
//         *res = append(*res, tmp)
//     }
//     used := make(map[int]bool, 0)

//     for i, val := range(nums) {
//         if !record[i] && !used[val] {
//             vector = append(vector, val)
//             used[val] = true
//             record[i] = true

//             helper(nums, vector, record, res)

//             record[i] = false
//             vector = vector[:curtLen]
//         }
//     }
// }

/**
** Solution2:
**/
func helper(nums, vector []int, record []bool, res *[][]int) {
	curtLen, n := len(vector), len(nums)
	if curtLen == len(nums) {
		tmp := make([]int, n)
		copy(tmp, vector)
		*res = append(*res, tmp)
	}

	for i := 0; i < n; i++ {
		if !record[i] && checkUsed(nums, record, i) {
			vector = append(vector, nums[i])
			record[i] = true

			helper(nums, vector, record, res)

			record[i] = false
			vector = vector[:curtLen]
		}
	}
}

// Function to check if the element can be used, if false it will create duplicated results
func checkUsed(nums []int, record []bool, index int) bool {
	if index > 0 {
		if nums[index] == nums[index-1] && !record[index-1] {
			return false
		}
	}
	return true
}
