package leet

import "testing"

func TestNumOfUnplacedFruits(t *testing.T) {
	type args struct {
		fruits  []int
		baskets []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example 1", args: args{fruits: []int{4, 2, 5}, baskets: []int{3, 5, 4}}, want: 1},
		{name: "example 2", args: args{fruits: []int{3, 6, 1}, baskets: []int{6, 4, 7}}, want: 0},
		{name: "example 3", args: args{fruits: []int{5}, baskets: []int{3}}, want: 1},
		{name: "example 4", args: args{fruits: []int{44, 10}, baskets: []int{26, 5}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumOfUnplacedFruits(tt.args.fruits, tt.args.baskets); got != tt.want {
				t.Errorf("NumOfUnplacedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
