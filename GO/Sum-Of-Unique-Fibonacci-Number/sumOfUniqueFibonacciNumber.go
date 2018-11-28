package sumOfUniqueFibonacciNumber

// Given positive integers n and k, determine if n can be expressed as a sum of k numbers in
// Fibonacci sequence (n <= 10^9, k <= 5)

// Example:

// Given n = 6, k = 2, you should return True.
// Given n = 5, k = 3, you should return False.
// Given n = 10059560, k = 4, you should return True.

func sumOfUniqueFibonacciNumber(n, k int) bool {
	array := getArray(n)
	return checkRes(array, n, k)
}

func getArray(n int) []int {
	array := make([]int, 0)
	s, e := 1, 1
	for e <= n {
		array = append(array, e)
		s, e = e, s+e
	}

	return array
}

func checkRes(array []int, n, k int) bool {
	arrayLen := len(array)
	if arrayLen == 0 {
		return false
	}
	if k == 1 {
		for element := range array {
			if element == n {
				return true
			}
		}
		return false
	}

	for i := arrayLen-1; i > 0; i-- {
		if checkRes(array[:i-1], n-array[i], k-1) {
			return true
		}
	}

	return false
}