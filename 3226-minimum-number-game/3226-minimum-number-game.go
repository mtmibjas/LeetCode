func numberGame(nums []int) []int {
    sort.Ints(nums)
    a := 0
    b := 1
    arr := []int{}
    for b < len(nums) {

        arr = append(arr, nums[b], nums[a])
        a += 2
        b += 2
    }

    return arr
}