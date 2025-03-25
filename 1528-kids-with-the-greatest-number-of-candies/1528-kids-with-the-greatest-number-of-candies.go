func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	arr := []bool{}
	for _, m := range candies {
		if max < m {
			max = m
		}
	}
	for _, m := range candies {
		if m+extraCandies >= max {
			arr = append(arr, true)
		} else {
			arr = append(arr, false)
		}
	}
	return arr

}