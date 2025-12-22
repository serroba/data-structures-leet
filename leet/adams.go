package leet

import (
	"slices"
	"strconv"
	"strings"
)

func PrintNumbers(list []int, maxNumber int) string {
	slices.Sort(list)

	var output []string

	length := len(list)
	if length == 0 {
		return "0-" + strconv.Itoa(maxNumber)
	}

	if list[0] == 1 {
		output = append(output, "0")
	} else if list[0] > 1 {
		output = append(output, "0-"+strconv.Itoa(list[0]-1))
	}

	for i := range list {
		output = append(output, strconv.Itoa(list[i]))
		if i == length-1 {
			break
		}

		if list[i+1]-list[i] == 2 {
			output = append(output, strconv.Itoa(list[i]+1))
		} else if list[i+1]-list[i] > 2 {
			output = append(output, strconv.Itoa(list[i]+1)+"-"+strconv.Itoa(list[i+1]-1))
		}
	}

	if maxNumber-list[length-1] == 1 {
		output = append(output, strconv.Itoa(maxNumber))
	} else if maxNumber-list[length-1] > 1 {
		output = append(output, strconv.Itoa(list[length-1]+1)+"-"+strconv.Itoa(maxNumber))
	}

	return strings.Join(output, ", ")
}
