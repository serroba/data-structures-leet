package leet

import "testing"

func TestMaxSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Add on empty list"},
		{name: "example 1", args: args{nums: []int{1, 2, 3, 4, 5}}, want: 15},
		{name: "example 2", args: args{nums: []int{1, 1, 0, 1, 1}}, want: 1},
		{name: "example 3", args: args{nums: []int{1, 2, -1, -2, 1, 0, -1}}, want: 3},
		{name: "example 4", args: args{nums: []int{2}}, want: 2},
		{name: "example 5", args: args{nums: []int{-20, 20}}, want: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSum(tt.args.nums); got != tt.want {
				t.Errorf("MaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
