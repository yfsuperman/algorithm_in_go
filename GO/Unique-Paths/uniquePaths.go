package uniquePaths

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

// The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

// How many possible unique paths are there?

// Note: m and n will be at most 100.

// Example 1:

// Input: m = 3, n = 2
// Output: 3
// Explanation:
// From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Right -> Down
// 2. Right -> Down -> Right
// 3. Down -> Right -> Right
// Example 2:

// Input: m = 7, n = 3
// Output: 28

func uniquePaths(m int, n int) int {
    record := make([]int, m)
    for i := 0; i < m; i++ {
        record[i] = 1
    }
    
    for j := 1; j < n; j ++ {
        for i := 0; i < m; i++ {
            if i > 0 {
                record[i] = record[i] + record[i - 1]
            } else {
                record[i] = 1
            }
            
        }
    }
    
    return record[m - 1]
}
