func separateDigits(nums []int) []int {
    arr := []int{}
    for i := 0; i < len(nums); i++{
        d := strconv.Itoa(nums[i])
        for _, n := range d {
            arr = append(arr, int(n - '0'))
        }   
    }
    return arr
}