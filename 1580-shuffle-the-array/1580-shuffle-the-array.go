func shuffle(nums []int, n int) []int {
    
    xi := 0
    yi := n
    var arr []int
    for yi < len(nums) {
        arr = append(arr, nums[xi], nums[yi])
        xi++
        yi++
    }
    return arr
}