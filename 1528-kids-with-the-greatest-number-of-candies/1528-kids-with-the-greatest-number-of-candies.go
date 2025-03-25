func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	arr := make([]bool, len(candies))
	for _, m := range candies {
		if max < m {
			max = m
		}
	}
	for i, m := range candies {
		arr[i] = m+extraCandies >= max
	}
	return arr

}