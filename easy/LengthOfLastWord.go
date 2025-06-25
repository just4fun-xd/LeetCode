package easy

func lengthOfLastWord(s string) int {
	n := 0
	for i := len(s) - 1; i >= 0; i-- {
		if n != 0 && string(s[i]) == " " {
			break
		} else if string(s[i]) != " " {
			n++
		}
	}
	return n
}
