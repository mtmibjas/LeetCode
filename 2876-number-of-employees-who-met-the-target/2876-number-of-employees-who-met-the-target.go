func numberOfEmployeesWhoMetTarget(hours []int, target int)(count int) {
    for i := 0; i < len(hours); i++{
        if hours[i] >= target{
            count++
        }
    }
    return
}