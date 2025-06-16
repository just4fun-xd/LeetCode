func isPalindromeNaive(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0

	for x > reversed {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}
	return x == reversed || x == reversed/10
}
