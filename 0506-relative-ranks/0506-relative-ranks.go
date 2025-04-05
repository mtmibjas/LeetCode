func findRelativeRanks(score []int) []string {
	temp := []int{}
    temp = append(temp, score...)
	sort.Slice(score, func(i, j int)bool{
        return score[i] > score[j] 
    })
	m := make(map[int]string)

	for i := 0; i < len(score); i++ {
		s := ""
		if i+1 == 1 {
			s = "Gold Medal"
		} else if i+1 == 2 {
			s = "Silver Medal"
		} else if i+1 == 3 {
			s = "Bronze Medal"
		} else {
			s = strconv.Itoa(i+1)
		}
		m[score[i]] = s
	}

	arr := []string{}
	for i := 0; i < len(temp); i++ {
		v, _ := m[temp[i]]
		arr = append(arr, v)
	}
	return arr
}