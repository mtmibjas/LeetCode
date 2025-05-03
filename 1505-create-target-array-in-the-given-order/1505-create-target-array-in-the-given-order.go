func createTargetArray(nums []int, index []int) []int {

    arr := make([]int, len(nums))
    for i, _ := range arr {
        arr[i] = -1
    }
    for i, v := range index {
        if arr[v] != -1 {
            for j := len(nums)-2; j >= v; j--{
                arr[j+1] = arr[j]
            }
        }
        arr[v] = nums[i]
    }

    return arr
}