func maxArea(height []int) int {

    r,l := 0, len(height)-1;
    max := 0
    for r < l {
        if m := min(height[r], height[l]) * (l - r); m > max {
            max = m
        }
        if height[r] > height[l]{
            l--
        }else{
            r++
        }
    }
    return max
}

func min(i,j int) int {
    if i > j {
        return j
    }
    return i
}