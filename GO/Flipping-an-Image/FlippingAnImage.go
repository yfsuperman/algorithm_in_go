package FlippingAnImage

// Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.

// To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

// To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].


func flipAndInvertImage(A [][]int) [][]int {    
    for _, a := range A {
        flipRow(a)
    }
    return A
}

func flipRow(row []int) {

    for i := 0; 2 * i <= len(row) - 1; i++ {
        if 2 * i < len(row) - 1 {
            row[i], row[len(row) - 1 - i] = 1- row[len(row) - 1 - i], 1- row[i]
        } else {
            row[i] = 1- row[i]
        }
    }
}