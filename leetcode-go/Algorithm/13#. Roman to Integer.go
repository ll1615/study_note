// Given a roman numeral, convert it to an integer.

// Input is guaranteed to be within the range from 1 to 3999.

package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MMCCCXCIX"))
}

func romanToInt(s string) int {
	ln := len(s)

	flag := []int{0, 0, 0, 0, 0, 0}

	var num int
	for i := ln - 1; i >= 0; i-- {
		switch s[i] {
		case 'I': //1
			if flag[0] == 0 {
				num++
			} else {
				num--
			}
		case 'V': //5
			flag[0] = 1
			if flag[1] == 0 {
				num += 5
			} else {
				num -= 5
			}
		case 'X': //10
			flag[0], flag[1] = 1, 1
			if flag[2] == 0 {
				num += 10
			} else {
				num -= 10
			}
		case 'L': //50
			flag[2] = 1
			if flag[3] == 0 {
				num += 50
			} else {
				num -= 50
			}
		case 'C': //100
			flag[2], flag[3] = 1, 1
			if flag[4] == 0 {
				num += 100
			} else {
				num -= 100
			}
		case 'D': //500
			flag[4] = 1
			if flag[5] == 0 {
				num += 500
			} else {
				num -= 500
			}
		case 'M': //1000
			flag[4], flag[5] = 1, 1
			num += 1000
		}
	}

	return num
}
//output: 2399
