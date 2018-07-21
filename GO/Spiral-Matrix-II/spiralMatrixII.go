package spiralMatrixII

// Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

// Example:

// Input: 3
// Output:
// [
//  [ 1, 2, 3 ],
//  [ 8, 9, 4 ],
//  [ 7, 6, 5 ]
// ]

func generateMatrix(n int) [][]int {
    var res = [][]int{}
    res = make([][]int, n)
    for i := range res {
        res[i] = make([]int, n)
    }
    directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    cnt, direct, row, col := 1, 0, 0, 0
    
    for cnt <= n*n {
        for row >= 0 && row < n && col >=0 && col < n && res[row][col] == 0 {
            res[row][col] = cnt
            cnt++
            row, col = row + directions[direct][0], col + directions[direct][1]
        }
        
        row, col = row - directions[direct][0], col - directions[direct][1]
        direct = (direct + 1)%4
        row, col = row + directions[direct][0], col + directions[direct][1]
    }
    
    return res
}