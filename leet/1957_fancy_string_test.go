package leet

import "testing"

func TestMakeFancyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example_1",
			args: args{s: "leeetcode"},
			want: "leetcode",
		},
		{
			name: "example_2",
			args: args{s: "aaabaaaa"},
			want: "aabaa",
		},
		{
			name: "example_3",
			args: args{s: "aab"},
			want: "aab",
		},
		{
			name: "single_char",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "two_same_chars_kept",
			args: args{s: "aa"},
			want: "aa",
		},
		{
			name: "three_same_chars_reduced_to_two",
			args: args{s: "aaa"},
			want: "aa",
		},
		{
			name: "four_same_chars_reduced_to_two",
			args: args{s: "aaaa"},
			want: "aa",
		},
		{
			name: "already_fancy_no_change",
			args: args{s: "abababab"},
			want: "abababab",
		},
		{
			name: "mixed_runs_multiple_deletions",
			args: args{s: "aabbbcccccdd"},
			want: "aabbccdd",
		},
		{
			name: "runs_separated_by_other_chars_do_not_interfere",
			args: args{s: "xxxyyyz"},
			want: "xxyyz",
		},
		{
			name: "no_three_consecutive_after_processing",
			args: args{s: "abbcccddddeeeee"},
			want: "abbccddee",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeFancyString(tt.args.s); got != tt.want {
				t.Errorf("MakeFancyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
