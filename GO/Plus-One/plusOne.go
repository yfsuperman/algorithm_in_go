package plusOne

// Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

// The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

// You may assume the integer does not contain any leading zero, except the number 0 itself.

// Example 1:

// Input: [1,2,3]
// Output: [1,2,4]
// Explanation: The array represents the integer 123.
// Example 2:

// Input: [4,3,2,1]
// Output: [4,3,2,2]
// Explanation: The array represents the integer 4321.

func plusOne(digits []int) []int {
    flag, l := 1, len(digits)
    for i := 1; i <= l; i++ {
        if flag == 0 {
            break
        }
        tmp := digits[l-i]
        digits[l-i], flag = (tmp + flag) % 10, (tmp + flag) / 10
    }
    
    if flag == 1 {
        res := []int{1}
        res = append(res, digits...)
        return res
    }
    
    return digits
}