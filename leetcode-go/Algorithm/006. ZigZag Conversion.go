func convert(s string, numRows int) string {
	length := len(s)
	if length <= numRows {
		return s
	}

	res := make([]string, numRows)
	gap := numRows - 2
	for i := 0; i < length; {
		for j := 0; j < numRows && i < length; j++ {
			res[j] += string(s[i])
			i++
		}
		for j := gap; j > 0 && i < length; j-- {
			res[j] += string(s[i])
			i++
		}
	}
	result := ""
	for i := 0; i < numRows; i++ {
		result += res[i]
	}
	return result
}
