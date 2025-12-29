package leet

func longestPalindrome(s string) string {
	if len(s) <= 2 && isPalindrome(s) {
		return s
	}

	longestLen := 0
	longestWord := ""

	for i := 0; i <= len(s)-1; i++ {
		for j := i; j <= len(s); j++ {
			if isPalindrome(s[i:j]) && len(s[i:j]) > longestLen {
				longestWord = s[i:j]
				longestLen = len(s[i:j])
			}
		}
	}

	return longestWord
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	for s[0] == s[len(s)-1] && len(s) > 2 {
		s = s[1 : len(s)-1]
	}

	if len(s) == 2 {
		return s[0] == s[1]
	}

	return len(s) <= 1
}
