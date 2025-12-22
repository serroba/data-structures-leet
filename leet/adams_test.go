package leet

import "testing"

func TestPrintNumbers(t *testing.T) {
	type args struct {
		list      []int
		maxNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test 1", args: args{list: []int{0, 1, 4, 6}, maxNumber: 10}, want: "0, 1, 2-3, 4, 5, 6, 7-10"},
		{name: "test 2", args: args{list: []int{3, 5}, maxNumber: 10}, want: "0-2, 3, 4, 5, 6-10"},
		{name: "test 2", args: args{list: []int{3, 10}, maxNumber: 10}, want: "0-2, 3, 4-9, 10"},
		{name: "test 2", args: args{list: []int{0, 1, 2}, maxNumber: 2}, want: "0, 1, 2"},
		{name: "test 2", args: args{list: []int{}, maxNumber: 5}, want: "0-5"},
		{name: "test 2", args: args{list: []int{0}, maxNumber: 0}, want: "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintNumbers(tt.args.list, tt.args.maxNumber); got != tt.want {
				t.Errorf("PrintNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
