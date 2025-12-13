package training

import (
	"slices"
)

var validBusinessLines = []string{"electronics", "grocery", "pharmacy", "restaurant"}

func ValidateCoupons(code []string, businessLine []string, isActive []bool) []string {
	validCoupons := map[string][]string{}
	for i := range code {
		if hasValidCharacters(code[i]) && inValidBusinessCategory(businessLine[i]) && isActive[i] {
			validCoupons[businessLine[i]] = append(validCoupons[businessLine[i]], code[i])
		}
	}
	var output []string
	for _, bl := range validBusinessLines {
		slices.Sort(validCoupons[bl])
		output = append(output, validCoupons[bl]...)
	}
	return output
}

func hasValidCharacters(code string) bool {
	if code == "" {
		return false
	}

	for _, r := range code {
		if !(r >= 'a' && r <= 'z' ||
			r >= 'A' && r <= 'Z' ||
			r >= '0' && r <= '9' ||
			r == '_') {
			return false
		}
	}
	return true
}

func inValidBusinessCategory(category string) bool {
	return slices.Contains(validBusinessLines, category)
}
