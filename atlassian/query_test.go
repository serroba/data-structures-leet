package atlassian_test

import (
	"reflect"
	"testing"

	"ds/atlassian"
)

func TestRetrieve(t *testing.T) {
	type args struct {
		query atlassian.Query
		items []atlassian.Item
	}
	tests := []struct {
		name string
		args args
		want []atlassian.Item
	}{
		{name: "Test on empty set", want: nil},
		{name: "Search by Project name", args: args{
			query: atlassian.Query{
				Operator: "&&",
				Project:  "some Project",
			},
			items: []atlassian.Item{
				{Project: "some Project"},
			},
		}, want: []atlassian.Item{
			{Project: "some Project"},
		}},
		{name: "Search by Project name and Project type", args: args{
			query: atlassian.Query{
				Operator: "&&",
				Project:  "some Project",
				Type:     atlassian.ItemType("bug"),
			},
			items: []atlassian.Item{
				{Project: "some Project", Kind: atlassian.ItemType("bug")},
			},
		}, want: []atlassian.Item{
			{Project: "some Project", Kind: atlassian.ItemType("bug")},
		}},
		{
			name: "Does not find by Project name and Project type", args: args{
				query: atlassian.Query{
					Operator: "&&",
					Project:  "some Project",
					Type:     atlassian.ItemType("bug"),
				},
				items: []atlassian.Item{
					{Project: "some Project", Kind: atlassian.ItemType("story")},
				},
			}, want: nil,
		},
		{
			name: "Find by Project name or Project type", args: args{
				query: atlassian.Query{
					Operator: "||",
					Project:  "some Project",
					Type:     atlassian.ItemType("story"),
				},
				items: []atlassian.Item{
					{Project: "some Project", Kind: atlassian.ItemType("story")},
				},
			}, want: []atlassian.Item{
				{Project: "some Project", Kind: atlassian.ItemType("story")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := atlassian.Retrieve(tt.args.query, tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}
