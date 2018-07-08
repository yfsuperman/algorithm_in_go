package HammingDistance

// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

// Given two integers x and y, calculate the Hamming distance.

// Note:
// 0 ≤ x, y < 231.

func hammingDistance(x int, y int) int {
    res := 0
    
    for (x > 0 || y > 0) {
        if (x ^ y) & 1 == 1 {
            res ++
        }
        x = x >> 1
        y = y >> 1
    }
    
    return res
}