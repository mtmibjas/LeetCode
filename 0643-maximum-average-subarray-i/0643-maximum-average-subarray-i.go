func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	mx := sum
	for i,j := 0,k; j < len(nums); i,j = i+1, j+1 {
		sum -= nums[i]
		sum += nums[j]
		if sum > mx {
			mx = sum
		}
	}

	return float64(mx) / float64(k)
}