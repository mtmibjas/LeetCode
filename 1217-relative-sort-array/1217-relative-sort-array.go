func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr := make([]int, 0)
    ars := make([]int, 0) 
	m := make(map[int]int)
    p := make(map[int]struct{}) 
    for i := 0; i < len(arr2); i++ {
		p[arr2[i]] = struct{}{}
	}
	for i := 0; i < len(arr1); i++ {
		m[arr1[i]]++
        if _, ok := p[arr1[i]]; !ok {
            ars = append(ars, arr1[i])
        }
	}
	for i := 0; i < len(arr2); i++ {
		v, _ := m[arr2[i]]
		for v > 0 {
			arr = append(arr, arr2[i])
			v--
		}
	}
	sort.Ints(ars)
    arr = append(arr, ars...)
	return arr
}