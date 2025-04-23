func finalValueAfterOperations(operations []string) int {
    sum := 0
    for _, v := range operations {
        if v == "X++" || v == "++X" {
            sum++
        } else{
            sum--
        }
    }

    return sum
}