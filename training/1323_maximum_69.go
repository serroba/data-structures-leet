package training

import "strconv"

func Maximum69Number(num int) int {
	numAsString := []rune(strconv.Itoa(num))
	for i, digit := range numAsString {
		if digit == '6' {
			numAsString[i] = '9'
			break
		}
	}
	newNum, _ := strconv.Atoi(string(numAsString))
	return newNum
}
