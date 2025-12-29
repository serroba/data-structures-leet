package leet

func isPalindromeString(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < len(s) && !isAlphaNumeric(s[i]) {
			i++
		}

		for j >= 0 && !isAlphaNumeric(s[j]) {
			j--
		}

		if i > j {
			return true
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
	}

	return true
}

func toLower(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return char + 'a' - 'A'
	}

	return char
}

func isAlphaNumeric(char byte) bool {
	return char >= 'a' && char <= 'z' ||
		char >= 'A' && char <= 'Z' ||
		char >= '0' && char <= '9'
}
