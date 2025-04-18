func numOfUnplacedFruits(fruits []int, baskets []int) int {
    c := 0
    for _, f := range fruits {
        is := false
        for  i := 0; i < len(baskets); i++ {
            if f <= baskets[i] {
                is = true
                baskets[i] = 0
                break
            }
        }    
        if !is {
            c++
        }
    }
    return c
}