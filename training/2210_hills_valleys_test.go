package training

import "testing"

func TestCountHillValley(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example 1", args: args{nums: []int{2, 4, 1, 1, 6, 5}}, want: 3},
		{name: "example 2", args: args{nums: []int{6, 6, 5, 5, 4, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountHillValley(tt.args.nums); got != tt.want {
				t.Errorf("CountHillValley() = %v, want %v", got, tt.want)
			}
		})
	}
}
