package nQueensII

// Given an integer n, return the number of distinct solutions to the n-queens puzzle.

// Example:

// Input: 4
// Output: 2
// Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
// [
//  [".Q..",  // Solution 1
//   "...Q",
//   "Q...",
//   "..Q."],

//  ["..Q.",  // Solution 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]

func totalNQueens(n int) int {
    res := 0
    queenPosition := make([]int, n)
    dfs(0, queenPosition, &res)
    return res
}

func dfs(row int, queenPosition []int, res *int) {
    for i := 0; i < len(queenPosition); i++ {
        if isValid(row, i, queenPosition) {
            queenPosition[row] = i
            if row == len(queenPosition) - 1 {
                *res ++
            }
            dfs(row+1, queenPosition, res)
        }
    }
}

func isValid(row, col int, queenPosition []int) bool {
    for i := 0; i < row; i++ {
        if col == queenPosition[i] || col - queenPosition[i] == row - i || col - queenPosition[i] == i - row {
            return false
        }
    }
    return true
}