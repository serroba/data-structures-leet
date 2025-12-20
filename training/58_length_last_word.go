package training

func lengthOfLastWord(s string) int {
	var lastWordLen, wordLen int
	for _, ch := range s {
		if ch == ' ' {
			if wordLen > 0 {
				lastWordLen = wordLen
			}
			wordLen = 0
		} else {
			wordLen++
		}
	}
	if wordLen > 0 {
		lastWordLen = wordLen
	}
	return lastWordLen
}
