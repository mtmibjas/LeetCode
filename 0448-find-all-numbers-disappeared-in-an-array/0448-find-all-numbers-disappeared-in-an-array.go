func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}
	arr := make([]int, 0)

	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			arr = append(arr, i)
		}
	}
	return arr
}