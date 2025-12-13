package training

func MakeFancyString(s string) string {
	if len(s) <= 2 {
		return s
	}

	output := make([]byte, 0, len(s))
	output = append(output, s[0], s[1])

	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i-1] && s[i-1] == s[i] {
			continue
		}
		output = append(output, s[i])
	}
	return string(output)
}
