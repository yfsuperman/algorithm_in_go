package spiralMatrix

// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

// Example 1:

// Input:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ]
// ]
// Output: [1,2,3,6,9,8,7,4,5]
// Example 2:

// Input:
// [
//   [1, 2, 3, 4],
//   [5, 6, 7, 8],
//   [9,10,11,12]
// ]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]

func spiralOrderSolution1(matrix [][]int) []int {
    // This method will create an extra 2D array to track if the position has been passed or not
	// If we can assume all the number in matrix is non-negtive, then we can mark matrix[r][c]=-1 to save nm space
	// This solution takes a lot of operation and will run out of time
	// spiralOrderSolution2 is more direct and cost less time
    
    m, n := len(matrix), len(matrix[0])
    res := []int{}
    var position = [][]int{}
    position = make([][]int, n)
    for i := range position {
        position[i] = make([]int, m)
    }
    directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    cnt, direct, row, col := 0, 0, 0, 0
    
    for cnt < n*n {
        for row >= 0 && row < n && col >= 0 && col < m && position[row][col] == 0 {
            res = append(res, matrix[row][col])
            position[row][col] = 1
            row, col, cnt = row + directions[direct][0], col + directions[direct][1], cnt+1
        }
        row, col = row - directions[direct][0], col - directions[direct][1]
        direct = (direct + 1) % 4
        row, col = row + directions[direct][0], col + directions[direct][1]
    }
    
    return res
}

func spiralOrderSolution2(matrix [][]int) []int {
    res := []int{}
    if matrix == nil || len(matrix) == 0 || len(matrix) == 0 { return res }
    
    l, r, t, b := 0, len(matrix[0]) - 1, 0, len(matrix) - 1
    
    for {
        for i := l; i <= r; i++ { res = append(res, matrix[t][i]) }
        t ++
        if t > b { break }
        
        for i := t; i <= b; i++ { res = append(res, matrix[i][r]) }
        r --
        if l > r { break }
        
        for i := r; i >= l; i-- { res = append(res, matrix[b][i]) }
        b --
        if t > b { break }
        
        for i := b; i >= t; i-- { res = append(res, matrix[i][l]) }
        l ++
        if l > r { break }
    }
    
    return res
}