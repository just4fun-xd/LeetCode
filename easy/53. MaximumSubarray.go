package easy

func maxSubArray(nums []int) int {
	curSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if curSum+nums[i] < nums[i] {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		if curSum > maxSum {
			maxSum = curSum

		}
	}
	return maxSum
}
