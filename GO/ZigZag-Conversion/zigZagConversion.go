package zigZagConversion

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Faster solution would be visiting the characters in the same order as reading the Zig-Zag pattern line by line.
// For all whole numbers kk,
// Characters in row 00 are located at indexes k \; (2 \cdot \text{numRows} - 2)k(2⋅numRows−2)
// Characters in row \text{numRows}-1numRows−1 are located at indexes k \; (2 \cdot \text{numRows} - 2) + \text{numRows} - 1k(2⋅numRows−2)+numRows−1
// Characters in inner row ii are located at indexes k \; (2 \cdot \text{numRows}-2)+ik(2⋅numRows−2)+i and (k+1)(2 \cdot \text{numRows}-2)- i(k+1)(2⋅numRows−2)−i.
func convert(s string, numRows int) string {
    if numRows == 1 { return s }
    
    var arr []string
    curtRow, row, down, res := 0, min(numRows, len(s)), 1, ""
    arr = make([]string, row)
    
    for _, char := range s {
        arr[curtRow] += string(char)
        curtRow += down
        if curtRow == 0 || curtRow == row - 1 {
            down = - down
        }
    }
    
    for _, str := range arr {
        res += str
    }
    
    return res
}

func min(a int, b int) int {
    if a > b { return b }
    return a
}