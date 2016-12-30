func longestCommonPrefix(strs []string) string {
	ln := len(strs)
	if ln < 1 {
		return ""
	}

	var lcp string
	var tmpChar byte

out:
	for i := 0; i < len(strs[0]); i++ {
		tmpChar = strs[0][i]
		for j := 0; j < ln; j++ {
			if i > len(strs[j])-1 || tmpChar != strs[j][i] {
				break out
			}
		}
		lcp += string(tmpChar)
	}
	return lcp
}
