func stableMountains(height []int, threshold int) []int {
    arr := []int{}
    for i := 1; i < len(height); i++{
        if height[i-1] > threshold {
            arr = append(arr, i)
        } 
    }

    return arr
}