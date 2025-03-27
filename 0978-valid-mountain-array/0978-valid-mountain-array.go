func validMountainArray(arr []int) bool {
    if len(arr) < 3 {
        return false
    }
    left := 0
    right := len(arr)-1
    move := right - left
    for left < right {
        if arr[left] == arr[left+1]|| arr[right] == arr[right-1] {
            return false
        }
        if arr[left] < arr[left+1] {
            left++
        }
      
        if arr[right] < arr[right-1] {
            right--
        }
        if move == right - left{
            return false
        }
        move = right -left
    }
    if left == 0 ||  right == len(arr)-1 {
        return false
    }
    return true
}