package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 4, -1, 7, 8}

	fmt.Println(maxSubArray(nums))

}

// 53. Maximum Subarray (Алгоритм Кадане)

func maxSubArray(nums []int) int {
	sum := nums[0]
	maxSum := nums[0]
	startTemp, startFin, endFin := 0, 0, 0

	for i := 1; i < len(nums); i++ {
		if sum+nums[i] < nums[i] {
			sum = nums[i]
			startTemp = i
		} else {
			sum += nums[i]
		}
		if sum > maxSum {
			maxSum = sum
			startFin = startTemp
			endFin = i
		}
	}
	fmt.Println(nums[startFin : endFin+1])
	return maxSum
}

/*
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
