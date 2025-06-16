func plusOne(digits []int) []int {
	k := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + k
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			k = 0
		}
	}
	return digits
}