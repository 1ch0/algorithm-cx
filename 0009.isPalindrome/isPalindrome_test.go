package _009_isPalindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is palindrome",
			args: args{x: 111},
			want: true,
		},
		{
			name: "is not palindrome",
			args: args{x: 211},
			want: false,
		},
		{
			name: "is not palindrome",
			args: args{x: 21},
			want: false,
		},
		{
			name: "is not palindrome",
			args: args{x: -11},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
