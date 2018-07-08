package TransposeMatrix

// Given a matrix A, return the transpose of A.

// The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

func transpose(A [][]int) [][]int {

	B := make([][]int, len(A[0]))
	
    for i:= 0; i < len(A[0]); i++ {
		B[i] = make([]int, len(A))
		
        for j:= 0; j< len(A); j++ {
            B[i][j] = A[j][i]
        }
    }
    
    return B
}