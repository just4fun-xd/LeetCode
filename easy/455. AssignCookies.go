package easy

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := len(g)-1, len(s)-1
	for i >= 0 && j >= 0 {
		if g[i] <= s[j] {
			i--
			j--
		} else {
			i--
		}

	}
	return len(s) - 1 - j
}
