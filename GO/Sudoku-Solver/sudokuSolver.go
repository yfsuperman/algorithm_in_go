package sudokuSolver

// Write a program to solve a Sudoku puzzle by filling the empty cells.

// A sudoku solution must satisfy all of the following rules:

// Each of the digits 1-9 must occur exactly once in each row.
// Each of the digits 1-9 must occur exactly once in each column.
// Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
// Empty cells are indicated by the character '.'.

// Note:

// The given board contain only digits 1-9 and the character '.'.
// You may assume that the given Sudoku puzzle will have a single unique solution.
// The given board size is always 9x9.

func solveSudoku(board [][]byte)  {
    solveSudokuHelper(board, 0, 0)
}

func solveSudokuHelper(board [][]byte, row, col int) bool {
    if col == 9 {
        return solveSudokuHelper(board, row + 1, 0)
    }
    if row == 9 {
        return true
    }
    if board[row][col] == '.' {
        for i := byte('1'); i <= '9'; i++ {
            board[row][col] = i
            if validSudoku(board, row, col) {
                if solveSudokuHelper(board, row, col + 1) {
                    return true
                }
            }
            board[row][col] = '.'
        }
    } else {
        return solveSudokuHelper(board, row, col + 1)
    }
    return false
}

func validSudoku(board [][]byte, row, col int) bool {
    current := board[row][col]
    for i := 0; i < 9; i++ {
        if i != col && board[row][i] == current {
            return false
        }
        if i != row && board[i][col] == current {
            return false
        }
    }
    
    subRow, subCol := row/3*3, col/3*3
    for i := subRow; i < subRow + 3; i++ {
        for j := subCol; j < subCol + 3; j ++ {
            if i != row && j != col && board[i][j] == current {
                return false
            }
        }
    }
    return true  
}