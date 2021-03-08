package main

import (
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"should be equal", args{"babad"}, "bab"},
	{"should be equal", args{"cbbd"}, "bb"},
	{"should be equal", args{"a"}, "a"},
	{"should be equal", args{"ac"}, "a"},
}

func Test_longestPalindrome(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindromeFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeFunc2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindromeFunc3(t *testing.T) {
	type args struct {
		s string
	}

	var tests = []struct {
		name string
		args args
		want string
	}{
		{"should be equal", args{"babad"}, "aba"},
		{"should be equal", args{"cbbd"}, "bb"},
		{"should be equal", args{"a"}, "a"},
		{"should be equal", args{"ac"}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeFunc3(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
