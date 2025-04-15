func differenceOfSum(nums []int) int {
    
    sum := 0
    total := 0
    for _, num := range nums {
        if num > 9 {
            sum += sumDigits(num)
        }else{
            sum += num
        }
        total += num
    }
    if total > sum {
        return total-sum
    }
    return sum - total
}

func sumDigits(num int) int {
    s := 0
	for num > 0 {
		s += num % 10
		num = num/10
	}
	

    return s
}