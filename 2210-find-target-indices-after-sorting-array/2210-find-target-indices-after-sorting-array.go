func targetIndices(nums []int, target int) []int {
    sort.Ints(nums)
    arr := make([]int, 0)

    for i, num := range nums{
        if num == target {
            arr = append(arr, i)
        }
    }
    return arr
}