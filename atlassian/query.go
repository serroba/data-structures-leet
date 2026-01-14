package atlassian

import (
	"slices"
	"strings"
)

type ItemType string

type Item struct {
	Key     string
	Project string
	Kind    ItemType
}

type Operator string

type Query struct {
	Operator Operator
	Project  string
	Type     ItemType
}

type Order string

type Sort struct {
	Type    Order
	Project Order
	// Key string
}

// project = ABC AND (kind = Bug OR kind = Story)
func Retrieve(query Query, sort *Sort, items ...Item) []Item {
	var result []Item
	for _, val := range items {
		switch query.Operator {
		case "&&":
			if val.Project == query.Project && val.Kind == query.Type {
				result = append(result, val)
			}

		case "||":
			if val.Project == query.Project || val.Kind == query.Type {
				result = append(result, val)
			}
		}
	}

	if sort != nil {
		if sort.Type == "ASC" {
			slices.SortFunc(result, func(i Item, j Item) int {
				// return strings.Compare(string(i.Kind), string(j.Kind))
			})
		}
	}

	return result
}
