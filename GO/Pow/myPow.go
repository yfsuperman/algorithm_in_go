package myPow

func myPow(x float64, n int) float64 {
    if n < 0 {
        return 1.0 / calcPow(x, -n)
    }

    return calcPow(x, n)
}

func calcPow(x float64, n int) float64 {
    if x == 0 {
        return 0
    }
    
    if n == 0 {
        return 1
    }
    
    res := calcPow(x, n >> 1)
    if n & 1 == 0 {
        return res * res
    }
    
    return res * res * x
}