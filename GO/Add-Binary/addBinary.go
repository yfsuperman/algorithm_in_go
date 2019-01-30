package addBinary

// Given two binary strings, return their sum (also a binary string).

// The input strings are both non-empty and contains only characters 1 or 0.

// Example 1:

// Input: a = "11", b = "1"
// Output: "100"
// Example 2:

// Input: a = "1010", b = "1011"
// Output: "10101"

func addBinary(a string, b string) string {
    arrA := []byte(a)
    arrB := []byte(b)
    max := findMax(len(arrA), len(arrB)) + 1
    res := make([]byte, max)
    flag := 0
    for i := 1; i <= max; i++ {
        valA, valB := 0, 0
        if len(arrA) - i >= 0 {
            valA = int(arrA[len(arrA) - i] - '0')
        }
        if len(arrB) - i >= 0 {
            valB = int(arrB[len(arrB) - i] - '0')
        }
        res[max - i] = byte((valA + valB + flag) % 2 + '0')
        flag = (valA + valB + flag) / 2
    }
    if res[0] == '0' {
        res = res[1:]
    }
    
    return string(res)
}

func findMax(a, b int) int {
    if a > b {
        return a
    }
    return b
}