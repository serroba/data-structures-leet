package training

import "testing"

func TestFindClosest(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test 1", args: args{2, 7, 4}, want: 1},
		{name: "test 2", args: args{2, 5, 6}, want: 2},
		{name: "test 3", args: args{1, 5, 3}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindClosest(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("FindClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
