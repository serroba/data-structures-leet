package training

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test on 0", args: args{n: 0}, want: false},
		{name: "test on 1", args: args{n: 1}, want: true},
		{name: "test on 3", args: args{n: 3}, want: true},
		{name: "test on 27", args: args{n: 27}, want: true},
		{name: "test on 5", args: args{n: 5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("IsPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
