package leet

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
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
		{name: "test 2", args: args{n: 2}, want: true},
		{name: "test 1024", args: args{n: 1024}, want: true},
		{name: "test 5", args: args{n: 5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("IsPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
