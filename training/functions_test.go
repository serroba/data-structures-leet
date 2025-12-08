package training_test

import (
	"reflect"
	"testing"

	"ds/training"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test 1", args: args{nums: []int{2, 7, 11, 15}, target: 9}, want: []int{0, 1}},
		{name: "test 2", args: args{nums: []int{3, 2, 4}, target: 6}, want: []int{1, 2}},
		{name: "test 3", args: args{nums: []int{3, 3}, target: 6}, want: []int{0, 1}},
		{name: "test 4", args: args{nums: []int{1, 5, 5, 2}, target: 10}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := training.TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *training.ListNode
		l2 *training.ListNode
	}
	tests := []struct {
		name string
		args args
		want *training.ListNode
	}{
		{name: "test 1", args: args{l1: nil, l2: nil}, want: nil},
		{
			name: "test 2",
			args: args{
				l1: &training.ListNode{
					Val: 2, Next: &training.ListNode{
						Val: 4, Next: &training.ListNode{
							Val: 3}}},
				l2: &training.ListNode{
					Val: 5, Next: &training.ListNode{
						Val: 6, Next: &training.ListNode{
							Val: 4}}}},
			want: &training.ListNode{
				Val: 7, Next: &training.ListNode{
					Val: 0, Next: &training.ListNode{
						Val: 8}}},
		},
		{
			name: "test 3",
			args: args{
				l1: &training.ListNode{
					Val: 9,
					Next: &training.ListNode{
						Val: 9,
						Next: &training.ListNode{
							Val: 9,
							Next: &training.ListNode{
								Val: 9,
								Next: &training.ListNode{
									Val: 9,
									Next: &training.ListNode{
										Val: 9,
										Next: &training.ListNode{
											Val: 9,
										},
									},
								},
							},
						},
					}},
				l2: &training.ListNode{
					Val: 9,
					Next: &training.ListNode{
						Val: 9,
						Next: &training.ListNode{
							Val: 9,
							Next: &training.ListNode{
								Val: 9,
							},
						},
					},
				},
			},
			want: &training.ListNode{
				Val: 8,
				Next: &training.ListNode{
					Val: 9,
					Next: &training.ListNode{
						Val: 9,
						Next: &training.ListNode{
							Val: 9,
							Next: &training.ListNode{
								Val: 0,
								Next: &training.ListNode{
									Val: 0,
									Next: &training.ListNode{
										Val: 0,
										Next: &training.ListNode{
											Val: 1,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := training.AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test 1", args: args{s: "abcabcbb"}, want: 3},
		{name: "test 2", args: args{s: "bbbbbb"}, want: 1},
		{name: "test 3", args: args{s: "pwwkew"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := training.LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test 0", args: args{x: 0}, want: true},
		{name: "test 1", args: args{x: 121}, want: true},
		{name: "test 10", args: args{x: 12321}, want: true},
		{name: "test 2", args: args{x: -121}, want: false},
		{name: "test 3", args: args{x: 10}, want: false},
		{name: "test 4", args: args{x: 5}, want: true},
		{name: "test 313", args: args{x: 313}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := training.IsPalindrome(tt.args.x); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRomanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test III", args: args{s: "III"}, want: 3},
		{name: "test LVIII", args: args{s: "LVIII"}, want: 58},
		{name: "test MCMXCIV", args: args{s: "MCMXCIV"}, want: 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := training.RomanToInt(tt.args.s); got != tt.want {
				t.Errorf("RomanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "example 1", args: args{strs: []string{"flower", "flow", "flight"}}, want: "fl"},
		{name: "example 2", args: args{strs: []string{"dog", "racecar", "car"}}, want: ""},
		{name: "example 2", args: args{strs: []string{"dog"}}, want: "dog"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := training.LongestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidParens(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "example 1", args: args{s: "()"}, want: true},
		{name: "example 2", args: args{s: "()[]{}"}, want: true},
		{name: "example 3", args: args{s: "(]"}, want: false},
		{name: "example 4", args: args{s: "([])"}, want: true},
		{name: "example 5", args: args{s: "([)]"}, want: false},
		{name: "example 6", args: args{s: "["}, want: false},
		{name: "example 7", args: args{s: "(("}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := training.IsValidParens(tt.args.s); got != tt.want {
				t.Errorf("IsValidParens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "example 1", args: args{nums: []int{1, 1, 2}}, want: []int{1, 2}},
		{name: "example 2", args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, want: []int{0, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := training.RemoveDuplicates(tt.args.nums); got != len(tt.want) {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, len(tt.want))
			}
			for i := range tt.want {
				if i < len(tt.want) && tt.args.nums[i] != tt.want[i] {
					t.Errorf("Not identical element = %v, want %v", tt.args.nums[i], tt.want[i])
				}
			}
		})
	}
}
