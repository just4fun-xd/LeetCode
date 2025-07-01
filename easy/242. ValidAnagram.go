package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := [26]int{}
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}
	for _, l := range freq {
		if l != 0 {
			return false
		}
	}
	return true
}
