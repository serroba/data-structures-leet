package leet

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test empty string", want: true},
		{name: "test single string", args: args{s: "a"}, want: true},
		{name: "test double string", args: args{s: "aa"}, want: true},
		{name: "test double string", args: args{s: "ab"}, want: false},
		{name: "test palindrome", args: args{s: "abcba"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "example 1", args: args{s: "babad"}, want: "bab"},
		{name: "example 2", args: args{s: "cbbd"}, want: "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
