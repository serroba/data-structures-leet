package leet

import "testing"

func Test_isPalindromeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// "A man, a plan, a canal: Panama" → true
		{
			name: "example 1 - palindrome with punctuation",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		// "race a car" → false
		{
			name: "example 2 - not a palindrome",
			args: args{s: "race a car"},
			want: false,
		},
		// " " → true (empty after removing non-alphanumeric)
		{
			name: "example 3 - space only",
			args: args{s: " "},
			want: true,
		},
		{
			name: "example 4 - non-alpha only",
			args: args{s: ".,"},
			want: true,
		},
		{
			name: "example 5 - with numbers",
			args: args{s: "0P"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeString(tt.args.s); got != tt.want {
				t.Errorf("isPalindromeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
