func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr := make([]int, 0)
    ars := make([]int, 0) 
	m := make(map[int]int)

	for i := 0; i < len(arr1); i++ {
		m[arr1[i]]++
        if !slices.Contains(arr2, arr1[i]){
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