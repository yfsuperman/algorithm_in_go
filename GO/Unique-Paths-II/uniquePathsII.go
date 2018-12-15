package uniquePathsII

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

// The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

// Now consider if some obstacles are added to the grids. How many unique paths would there be?

// An obstacle and empty space is marked as 1 and 0 respectively in the grid.

// Note: m and n will be at most 100.

// Example 1:

// Input:
// [
//   [0,0,0],
//   [0,1,0],
//   [0,0,0]
// ]
// Output: 2
// Explanation:
// There is one obstacle in the middle of the 3x3 grid above.
// There are two ways to reach the bottom-right corner:
// 1. Right -> Right -> Down -> Down
// 2. Down -> Down -> Right -> Right

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid) == 0  {
        return 0
    }
    if obstacleGrid[0][0] == 1 {
        return 0
    }
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    if obstacleGrid[m-1][n-1] == 1 {
        return 0
    }
    blocker := false
    for i := 0; i < m; i++ {
        if obstacleGrid[i][0] != 1 && !blocker {
            obstacleGrid[i][0] = 1
        } else {
            obstacleGrid[i][0] = 0
            blocker = true
        }
    }
    blocker = false
    for j := 1; j < n; j++ {
        if obstacleGrid[0][j] != 1 && !blocker {
            obstacleGrid[0][j] = 1
        } else {
            obstacleGrid[0][j] = 0
            blocker = true
        }
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                obstacleGrid[i][j] = 0
            } else {
                obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
            }
        }
    }
    
    return obstacleGrid[m-1][n-1]
    
}