package leet

import "testing"

func TestMaximum69Number(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test 6", args: args{num: 6}, want: 9},
		{name: "test 9", args: args{num: 9}, want: 9},
		{name: "test 69", args: args{num: 69}, want: 99},
		{name: "test 9669", args: args{num: 9669}, want: 9969},
		{name: "test 9996", args: args{num: 9996}, want: 9999},
		{name: "test 9999", args: args{num: 9999}, want: 9999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maximum69Number(tt.args.num); got != tt.want {
				t.Errorf("Maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
