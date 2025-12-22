package leet

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test 0", args: args{n: 0}, want: false},
		{name: "test 1", args: args{n: 1}, want: true},
		{name: "test 2", args: args{n: 2}, want: false},
		{name: "test 4", args: args{n: 4}, want: true},
		{name: "test 16", args: args{n: 16}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("IsPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
