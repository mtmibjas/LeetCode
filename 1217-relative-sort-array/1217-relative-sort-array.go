func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr := make([]int, len(arr1))
	m := make(map[int]int)
    index := len(arr2)
	for i := 0; i < len(arr1); i++ {
		m[arr1[i]]++
        if !slices.Contains(arr2, arr1[i]){
            arr[index] =  arr1[i]
            index++
        }
	}
    sort.Ints(arr[len(arr2):])
    index = 0
	for i := 0; i < len(arr2); i++ {
		v, _ := m[arr2[i]]
		for v > 0 {
			arr[index] =  arr2[i]
            index++
			v--
		}
	}
	
	return arr
}