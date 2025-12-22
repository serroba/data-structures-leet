package leet

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test 1", args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 3}, want: 2},
		{name: "test 1", args: args{nums: []int{1, 2, 2, 4, 5, 6, 7, 8, 9}, target: 3}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
