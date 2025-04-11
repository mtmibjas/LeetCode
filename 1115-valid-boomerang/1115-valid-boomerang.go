func isBoomerang(points [][]int) bool {
    return float64(0) != Area(points[0][0],points[0][1],points[1][0],points[1][1],points[2][0],points[2][1])
}


func Area(x1,y1,x2,y2,x3,y3 int) float64 {
    return 0.5 * float64(x1 * (y2 - y3) + x2 * (y3 - y1) + x3 * (y1 - y2))
}