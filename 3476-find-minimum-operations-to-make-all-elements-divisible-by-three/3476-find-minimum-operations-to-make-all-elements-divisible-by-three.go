func minimumOperations(nums []int) int {
    c := 0
    for _, num := range nums {
        if num%3 != 0 {
            c++
        }
    }
    return c
}