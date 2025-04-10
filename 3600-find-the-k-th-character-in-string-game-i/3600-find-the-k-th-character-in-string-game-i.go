func kthCharacter(k int) byte {
    k--
    //fmt.Println(k)
    count := 0
    for k > 0 {
        count++
       //fmt.Println(count)
        k &= k - 1
       // fmt.Println(k) 
    }
    
    return 'a' + byte(count%26)
}
