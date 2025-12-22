package leet

import (
	"slices"
	"testing"
)

func TestValidateCoupons(t *testing.T) {
	type args struct {
		code         []string
		businessLine []string
		isActive     []bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example_1",
			args: args{
				code:         []string{"SAVE20", "", "PHARMA5", "SAVE@20"},
				businessLine: []string{"restaurant", "grocery", "pharmacy", "restaurant"},
				isActive:     []bool{true, true, true, true},
			},
			want: []string{"PHARMA5", "SAVE20"},
		},
		{
			name: "example_2",
			args: args{
				code:         []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"},
				businessLine: []string{"grocery", "electronics", "invalid"},
				isActive:     []bool{false, true, true},
			},
			want: []string{"ELECTRONICS_50"},
		},
		{
			name: "sorts_by_business_line_then_lexicographically",
			args: args{
				code: []string{
					"b",
					"a",
					"zz",
					"aa",
					"p2",
					"p1",
					"r2",
					"r1",
				},
				businessLine: []string{
					"electronics",
					"electronics",
					"grocery",
					"grocery",
					"pharmacy",
					"pharmacy",
					"restaurant",
					"restaurant",
				},
				isActive: []bool{true, true, true, true, true, true, true, true},
			},
			want: []string{"a", "b", "aa", "zz", "p1", "p2", "r1", "r2"},
		},
		{
			name: "filters_out_inactive_even_if_other_fields_valid",
			args: args{
				code:         []string{"E1", "G1"},
				businessLine: []string{"electronics", "grocery"},
				isActive:     []bool{false, true},
			},
			want: []string{"G1"},
		},
		{
			name: "filters_out_invalid_business_line",
			args: args{
				code:         []string{"OK", "OK2"},
				businessLine: []string{"restaurant", "restaurants"},
				isActive:     []bool{true, true},
			},
			want: []string{"OK"},
		},
		{
			name: "validates_code_charset_allows_underscore_and_alnum_only",
			args: args{
				code:         []string{"A_B_9", "bad-hyphen", "bad space", "bad.dot", ""},
				businessLine: []string{"electronics", "electronics", "electronics", "electronics", "electronics"},
				isActive:     []bool{true, true, true, true, true},
			},
			want: []string{"A_B_9"},
		},
		{
			name: "all_invalid_returns_empty",
			args: args{
				code:         []string{"", "@@@", "OK"},
				businessLine: []string{"electronics", "grocery", "pharmacy"},
				isActive:     []bool{true, true, false},
			},
			want: []string{},
		},
		{
			name: "case_sensitivity_business_line_must_match_exactly",
			args: args{
				code:         []string{"E1", "E2"},
				businessLine: []string{"Electronics", "electronics"},
				isActive:     []bool{true, true},
			},
			want: []string{"E2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateCoupons(tt.args.code, tt.args.businessLine, tt.args.isActive); !slices.Equal(got, tt.want) {
				t.Errorf("ValidateCoupons() = %v, want %v", got, tt.want)
			}
		})
	}
}
