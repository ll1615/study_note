//method one
func reverse(x int) int {
	abx := x
	if x < 0 {
		abx = -x
	}

	revx := int64(0)
	for ; abx > 0; abx /= 10 {
		revx = revx*10 + int64(abx)%10
	}

	if x < 0 {
		revx = -revx
	}

	if revx > 0x7fffffff || revx < -0x7fffffff {
		return 0
	}

	return int(revx)
}

//method two
// func reverse(x int) int {
//     abx := x
// 	if x < 0 {
// 		abx = -x
// 	}

// 	strx := strconv.Itoa(abx)
// 	length := len(strx)
// 	strvx := ""

// 	for i := 0; i < length; i++ {
// 		strvx += string(strx[length-1-i])
// 	}

// 	revx, _ := strconv.Atoi(strvx)

// 	if x < 0 {
// 		revx = -revx
// 	}
	
// 	if revx > (1<<31-1) || revx < -(1<<31-1){
// 	    revx = 0
// 	}

// 	return revx
// }
