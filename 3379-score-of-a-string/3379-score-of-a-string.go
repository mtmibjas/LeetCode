func scoreOfString(s string) int {
    sum := 0
    for i := 1; i < len(s); i++ {
        sum += Abs(int(s[i-1]), int(s[i]))
    }
    return sum
}

func Abs(a, b int) int {

    if a-b < 0 {
        return b-a
    }
    return a-b
}