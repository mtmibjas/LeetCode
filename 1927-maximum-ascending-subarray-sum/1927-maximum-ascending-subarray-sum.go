func maxAscendingSum(nums []int) int {
	n := len(nums)
	maxSum := nums[0]
	currSum := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			currSum += nums[i]
		} else {
			if currSum > maxSum {
				maxSum = currSum
			}
			currSum = nums[i]
		}
	}
	if currSum > maxSum {
		maxSum = currSum
	}

	return maxSum
}
