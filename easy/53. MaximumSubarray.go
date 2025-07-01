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

/* Алгоритм Кадане
func maxSubArray(nums []int) int {
	curSum := nums[0]
	maxSum := nums[0]
	// startTemp, startFin, endFin  := 0, 0, 0

	for i := 1; i < len(nums); i++ {
		if curSum+nums[i] < nums[i] {
			curSum = nums[i]
			// startTemp = i + 1
		} else {
			curSum += nums[i]
		}
		if curSum > maxSum {
			maxSum = curSum
			// startFin = startTemp
			// endFin = i
		}
	}
	// fmt.Println(nums[startFin : endFin+1])
	// fmt.Printf("max_sum = %d, from  %d to %d\n", maxSum, startFin, endFin)
	return maxSum
} */
