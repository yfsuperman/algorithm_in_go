package multiplyStrings

// Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

// Example 1:

// Input: num1 = "2", num2 = "3"
// Output: "6"
// Example 2:

// Input: num1 = "123", num2 = "456"
// Output: "56088"
// Note:

// The length of both num1 and num2 is < 110.
// Both num1 and num2 contain only digits 0-9.
// Both num1 and num2 do not contain any leading zero, except the number 0 itself.
// You must not use any built-in BigInteger library or convert the inputs to integer directly.

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    
    numArr1 := []byte(num1)
    numArr2 := []byte(num2)
    tmp := make([]int, len(numArr1) + len(numArr2))
    
    for i := len(numArr1) - 1; i >= 0; i-- {
        for j := len(numArr2) - 1; j >= 0; j-- {
            tmp[i+j+1] += int(numArr1[i] - '0') * int(numArr2[j] - '0')
        }
    }
    
    for i := len(tmp) - 1; i > 0; i-- {
        tmp[i-1] += tmp[i] / 10
        tmp[i] = tmp[i] % 10
    }
    
    if tmp[0] == 0 {
        tmp = tmp[1:]
    }
    
    res := make([]byte, len(tmp))
	for i := 0; i < len(tmp); i++ {
		res[i] = '0' + byte(tmp[i])
	}

    return string(res)
}