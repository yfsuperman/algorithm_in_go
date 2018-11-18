package divideTwoIntegers

// Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

// Return the quotient after dividing dividend by divisor.

// The integer division should truncate toward zero.

// Example 1:

// Input: dividend = 10, divisor = 3
// Output: 3
// Example 2:

// Input: dividend = 7, divisor = -3
// Output: -2
// Note:

// Both dividend and divisor will be 32-bit signed integers.
// The divisor will never be 0.
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.

const (
	INT_MAX = 1<<31 - 1
	INT_MIN = -1 << 31
)
func divide(dividend int, divisor int) int {
    if divisor == 0 {
        return INT_MAX
    }
    
    absM, absN, resSign := analysis(dividend, divisor)
    res := 0
    if absN == 1 {
        res = absM
    } else {
        res = findRes(absM, absN)
    }
    if resSign == -1 {
        res = -res
    }
    if res > INT_MAX || res < INT_MIN {
        return INT_MAX
    }
    return res
}

func analysis(m, n int) (absM, absN, resSign int) {
    absM, absN, resSign = m, n, 1
    if m < 0 {
        absM = -m
        resSign = -resSign
    }
    if n < 0 {
        absN = -n
        resSign = -resSign
    }
    return absM, absN, resSign
}

func findRes(m, n int) int {
    res := 0
    for m >= n {
        tmp, sum := 1, n
        for m >= sum << 1 {
            sum <<= 1
            tmp <<= 1
        }
        res += tmp
        m -= sum
    }
    
    return res
}