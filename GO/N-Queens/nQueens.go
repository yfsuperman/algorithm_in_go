package nQueens

// Given an integer n, return all distinct solutions to the n-queens puzzle.

// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

// Example:

// Input: 4
// Output: [
//  [".Q..",  // Solution 1
//   "...Q",
//   "Q...",
//   "..Q."],

//  ["..Q.",  // Solution 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]
// Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.

func solveNQueens(n int) [][]string {
    res := make([][]string, 0)
    if n <= 0 {
        res = append(res, []string{})
        return res
    }
    queenPosition := make([]int, n)
    dfs(0, queenPosition, &res)
    
    return res
}

func dfs(row int, queenPosition []int, res *[][]string) {
    for i := 0; i < len(queenPosition); i ++ {
        if isValid(row, i, queenPosition) {
            queenPosition[row] = i
            if row == len(queenPosition) - 1 {
                tmp := constructResult(queenPosition)
                *res = append(*res, tmp)
            }
            dfs(row+1, queenPosition, res)
        }
    }
}

func isValid(row, col int, queenPosition []int) bool {
    for i := 0; i < row; i++ {
        diag := false
        if col - queenPosition[i] == row - i || col - queenPosition[i] == i - row {
            diag = true
        }
        if col == queenPosition[i] || diag {
            return false
        }
    }
    return true
}

func constructResult(queenPosition []int) []string {
    n := len(queenPosition)
    res := make([]string, 0)
    for i := 0; i < n; i++ {
        str := make([]byte, n)
        for j := 0; j < n; j++ {
            str[j] = '.'
        }
        str[queenPosition[i]] = 'Q'
        res = append(res, string(str))
    }
    return res
}