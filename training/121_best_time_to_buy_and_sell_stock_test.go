package training

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// [7,1,5,3,6,4] → 5 (buy at 1, sell at 6)
		{
			name: "example 1",
			args: args{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		// [7,6,4,3,1] → 0 (no profit possible, prices only decrease)
		{
			name: "example 2 - decreasing prices",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		// [1] → 0 (single day, can't buy and sell)
		{
			name: "example 3 - single price",
			args: args{prices: []int{1}},
			want: 0,
		},
		// [2,4,1] → 2 (buy at 2, sell at 4)
		{
			name: "example 4 - max profit before min price",
			args: args{prices: []int{2, 4, 1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
