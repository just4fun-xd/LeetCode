func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	for i := 0; i < minLen; i++ {
		ch := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if ch != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLen]
}