func numOfUnplacedFruits(fruits []int, baskets []int) int {
    for _, f := range fruits {
        for  i := 0; i < len(baskets); i++ {
            if f <= baskets[i] {
                baskets = append(baskets[:i],baskets[i+1:]...)
                break
            }
        }    
    }
    return len(baskets)
}