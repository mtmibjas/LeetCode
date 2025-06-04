func climbStairs(n int) int {

    if n <= 2 {
        return n
    }
    f := generatefib()
	f1,f2 := 0, 0 
    for i := 0; i <=n; i++ {
        if i == n  {
            f1 = f()
        }else if i == n-1 {
            f2 = f()
        }else{
            f()
        }
	}
    return f1+f2
}

func generatefib() func() int {
	f, s := 0, 1
	return func() int {
		n := f
		f, s = s, f+s
		return n
	}
}
