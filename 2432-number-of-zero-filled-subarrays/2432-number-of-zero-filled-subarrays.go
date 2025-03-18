func zeroFilledSubarray(nums []int) int64 {

	count := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
			continue
		}
		if count != 0 {
			sum += ((count + count*count) / 2)
		}
		count = 0

	}
	if count != 0 {

		sum += ((count + count*count) / 2)
	}

	return int64(sum)
}