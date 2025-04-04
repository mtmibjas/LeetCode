func divide(dividend int, divisor int) int {
    n := int(dividend/divisor)
    p := int(math.Pow(2, 31))
    if n > p-1 {
        n = p-1
    }
    if n < p*(-1) {
        n = p*(-1)
    }
    return n
}