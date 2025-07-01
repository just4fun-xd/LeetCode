package easy

func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	pos := len(nums) - 1
	newNums := make([]int, len(nums))

	for l <= r {
		leftSq := nums[l] * nums[l]
		rightSq := nums[r] * nums[r]
		if leftSq > rightSq {
			newNums[pos] = leftSq
			l++
		} else {
			newNums[pos] = rightSq
			r--
		}
		pos--
	}
	return newNums
}
