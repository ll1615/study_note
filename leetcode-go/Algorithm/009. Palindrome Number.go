func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}

	sx := ""
	for ; x != 0; x /= 10 {
		sx += string(x % 10)
	}
	ln := len(sx)

	for i := 0; i < (ln+1)/2; i++ {
		if sx[i] != sx[ln-i-1] {
			return false
		}
	}
	return true
}
